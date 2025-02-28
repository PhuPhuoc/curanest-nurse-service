package nurserepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
	"github.com/jmoiron/sqlx"
)

func (repo *nurseRepo) GetStaffByIds(ctx context.Context, ids nursequeries.StaffIdsQueryDTO) ([]nursedomain.Nurse, error) {
	if len(ids.Ids) == 0 {
		return nil, nil
	}

	dtos := []NurseDTO{}
	where := " where id in (?)"
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE, FIELD, nil) + where

	query, args, err := sqlx.In(query, ids.Ids)
	if err != nil {
		return nil, err
	}

	if err := repo.db.Select(&dtos, query, args...); err != nil {
		return nil, err
	}

	entities := make([]nursedomain.Nurse, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}
	return entities, nil
}
