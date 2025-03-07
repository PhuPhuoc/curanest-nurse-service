package nursequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

type getNurseWithFilterHandler struct {
	queryRepo NurseQueryRepo
}

func NewGetNursesWithFilterHandler(queryRepo NurseQueryRepo) *getNurseWithFilterHandler {
	return &getNurseWithFilterHandler{
		queryRepo: queryRepo,
	}
}

type NurseRequestQueryDTO struct {
	Filter *NurseFilterDTO `json:"filter"`
	Paging *common.Paging  `json:"paging"`
}

type NurseFilterDTO struct {
	ServiceId string  `json:"service-id"`
	NurseName string  `json:"nurse-name"`
	Rate      float64 `json:"rate"`
}

func (h *getNurseWithFilterHandler) Handle(ctx context.Context, rq *NurseRequestQueryDTO) ([]GetNurseDTO, error) {
	rq.Paging.Process()
	entities, err := h.queryRepo.GetByFilter(ctx, rq)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get nursing list").WithInner(err.Error())
	}

	dtos := make([]GetNurseDTO, len(entities))
	for i := range entities {
		dtos[i] = *toGetNurseDTO(&entities[i])
	}

	return dtos, nil
}
