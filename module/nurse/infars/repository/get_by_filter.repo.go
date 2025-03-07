package nurserepository

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
)

func (repo *nurseRepo) GetByFilter(ctx context.Context, requestQuery *nursequeries.NurseRequestQueryDTO) ([]nursedomain.Nurse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	whereClause, args := prepareWhereClause(requestQuery.Filter)

	errchan := make(chan error, 2)
	countchan := make(chan int, 1)
	datachan := make(chan []nursedomain.Nurse, 1)

	var wg sync.WaitGroup
	wg.Add(2)
	go repo.getCount(ctx, whereClause, args, errchan, countchan, &wg)
	go repo.getData(ctx, *requestQuery.Paging, whereClause, args, errchan, datachan, &wg)

	var once sync.Once // Make sure to close the channel only once.
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errchan <- fmt.Errorf("panic in goroutine: %v", r)
			}
		}()

		wg.Wait()
		once.Do(func() {
			close(errchan)
			close(countchan)
			close(datachan)
		})
	}()

	var totalRecord int
	var nursing []nursedomain.Nurse

	receivedCount := 0
	expectedCount := 2

	for {
		select {
		case err, ok := <-errchan:
			if ok {
				return nil, err
			}
		case count, ok := <-countchan:
			if ok {
				totalRecord = count
				receivedCount++
			}
		case data, ok := <-datachan:
			if ok {
				nursing = data
				receivedCount++
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("operation timed out: %w", ctx.Err())
		}

		if receivedCount == expectedCount {
			break
		}
	}

	totalPages := int(math.Ceil(float64(totalRecord) / float64(requestQuery.Paging.Size)))
	requestQuery.Paging.Total = totalPages

	return nursing, nil
}

func (r *nurseRepo) getData(
	ctx context.Context,
	paging common.Paging,
	whereClause string,
	args []interface{},
	errchan chan<- error,
	datachan chan<- []nursedomain.Nurse,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	args = append(args, paging.Size)
	args = append(args, (paging.Page-1)*paging.Size)
	limit := " limit ? offset ?"

	order := " order by nurse_name desc"

	fields := strings.Join(FIELD, ", ")

	join := ` join nursing_service on nursing_id = id`
	query := "select distinct " + fields + " from " + TABLE + join + whereClause + order + limit
	fmt.Println("query: ", query)
	var listdto []NurseDTO
	if err := r.db.SelectContext(ctx, &listdto, query, args...); err != nil {
		errchan <- err
		return
	}

	var entities []nursedomain.Nurse
	for _, dto := range listdto {
		entity, _ := dto.ToEntity()
		entities = append(entities, *entity)
	}

	datachan <- entities
}

func (r *nurseRepo) getCount(
	ctx context.Context,
	whereClause string,
	args []interface{},
	errchan chan<- error,
	countchan chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	join := ` join nursing_service on nursing_id = id`
	query := "select count(distinct id) from " + TABLE + join + whereClause

	var total int
	if err := r.db.GetContext(ctx, &total, query, args...); err != nil {
		errchan <- err
		return
	}

	countchan <- total
}

func prepareWhereClause(param *nursequeries.NurseFilterDTO) (string, []interface{}) {
	conditions := []string{}
	args := []interface{}{}

	if param.ServiceId != "" {
		conditions = append(conditions, "service_id = ?")
		args = append(args, param.ServiceId)
	}

	if param.NurseName != "" {
		conditions = append(conditions, "nurse_name like ?")
		args = append(args, "%"+param.NurseName+"%")
	}

	if param.Rate != 0 {
		conditions = append(conditions, "rate >= ?")
		args = append(args, strconv.FormatFloat(param.Rate, 'f', -1, 64))
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " where " + strings.Join(conditions, " and ")
	}

	return whereClause, args
}
