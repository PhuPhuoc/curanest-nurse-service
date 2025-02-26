package nursecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type mapNurseServiceHandler struct {
	cmdRepo NurseCommandRepo
}

func NewMapNurseServiceHandler(cmdRepo NurseCommandRepo) *mapNurseServiceHandler {
	return &mapNurseServiceHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *mapNurseServiceHandler) Handle(ctx context.Context, nurseId uuid.UUID, serviceIds []uuid.UUID) error {
	if err := h.cmdRepo.MapNurseService(ctx, nurseId, serviceIds); err != nil {
		return common.NewInternalServerError().
			WithReason("unable to map service with nursing").
			WithInner(err.Error())
	}
	return nil
}
