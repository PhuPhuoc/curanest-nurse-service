package nursequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type getNursingServiceHandler struct {
	queryRepo NurseQueryRepo
}

func NewGetNursingServiceHandler(queryRepo NurseQueryRepo) *getNursingServiceHandler {
	return &getNursingServiceHandler{
		queryRepo: queryRepo,
	}
}

func (h *getNursingServiceHandler) Handle(ctx context.Context, nursingId uuid.UUID) (*NursingServiceDTO, error) {
	serIds, err := h.queryRepo.GetNurseService(ctx, nursingId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list service-id of this nursing").WithInner(err.Error())
	}

	dto := &NursingServiceDTO{
		ServiceIds: serIds,
	}
	return dto, nil
}
