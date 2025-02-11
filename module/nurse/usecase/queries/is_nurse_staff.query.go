package nursequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type isNurseStaffHandler struct {
	queryRepo NurseQueryRepo
}

func NewIsNurseStaffHandler(queryRepo NurseQueryRepo) *isNurseStaffHandler {
	return &isNurseStaffHandler{
		queryRepo: queryRepo,
	}
}

func (h *isNurseStaffHandler) Handle(ctx context.Context, nurseId uuid.UUID) error {
	existed, err := h.queryRepo.IsStaffExisted(ctx, nurseId)
	if err != nil {
		return common.NewInternalServerError().
			WithReason("can't check if this nurse is a staff member or not").
			WithInner(err.Error())
	}

	if existed {
		return common.NewBadRequestError().WithReason("this nurse is already a staff member")
	}
	return nil
}
