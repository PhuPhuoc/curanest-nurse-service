package nursecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type updateNurseToStaffHandler struct {
	cmdRepo    NurseCommandRepo
	accService ExternalAccountService
}

func NewUpdateNurseToStaffHandler(cmdRepo NurseCommandRepo, accService ExternalAccountService) *updateNurseToStaffHandler {
	return &updateNurseToStaffHandler{
		cmdRepo:    cmdRepo,
		accService: accService,
	}
}

type UpdateRoleRequestRPC struct {
	Role string `json:"role"`
}

func (h *updateNurseToStaffHandler) Handle(ctx context.Context, nurseId uuid.UUID, categoryId uuid.UUID) error {
	if err := h.cmdRepo.UpdateNurseToStaff(ctx, nurseId, categoryId); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot convert nurse to staff").WithInner(err.Error())
	}

	payloadRPC := &UpdateRoleRequestRPC{
		Role: "staff",
	}

	if err := h.accService.UpdateRoleNurseToStaffRPC(ctx, nurseId, payloadRPC); err != nil {
		_ = h.cmdRepo.UpdateStaffToNurse(ctx, nurseId)
		return common.NewInternalServerError().
			WithReason("cannot update this nurse account to staff")
	}

	return nil
}
