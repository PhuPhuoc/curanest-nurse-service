package majorqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

type getMajorsHandler struct {
	queryRepo MajorQueryRepo
}

func NewGetMajorHandler(queryRepo MajorQueryRepo) *getMajorsHandler {
	return &getMajorsHandler{
		queryRepo: queryRepo,
	}
}

func (h *getMajorsHandler) Handle(ctx context.Context) ([]MajorDTO, error) {
	entities, err := h.queryRepo.GetAllMajors(ctx)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithMessage("cannot get list major").
			WithReason("error at GetAllMajors-repo").
			WithInner(err.Error())
	}

	list_dto := make([]MajorDTO, len(entities))
	for i := range entities {
		list_dto[i] = toDTO(&entities[i])
	}
	return list_dto, nil
}
