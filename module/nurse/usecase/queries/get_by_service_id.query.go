package nursequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type getNurseByServiceIdHandler struct {
	queryRepo NurseQueryRepo
}

func NewGetNurseByServiceIdHandler(queryRepo NurseQueryRepo) *getNurseByServiceIdHandler {
	return &getNurseByServiceIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getNurseByServiceIdHandler) Handle(ctx context.Context, serviceId uuid.UUID) ([]GetNurseDTO, error) {
	entities, err := h.queryRepo.GetByServiceId(ctx, serviceId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list nursing by serviceId (" + serviceId.String() + ")").WithInner(err.Error())
	}

	dtos := make([]GetNurseDTO, len(entities))
	for i := range entities {
		dtos[i] = *toGetNurseDTO(&entities[i])
	}

	return dtos, nil
}
