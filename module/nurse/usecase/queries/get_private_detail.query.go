package nursequeries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type getNursingPrivateDetailHandler struct {
	queryRepo NurseQueryRepo
	accRPC    ExternalAccountService
}

func NewGetNursingPrivateDetailHandler(queryRepo NurseQueryRepo, accRPC ExternalAccountService) *getNursingPrivateDetailHandler {
	return &getNursingPrivateDetailHandler{
		queryRepo: queryRepo,
		accRPC:    accRPC,
	}
}

func (h *getNursingPrivateDetailHandler) Handle(ctx context.Context, nurseId uuid.UUID) (*ResponseProfileDTO, error) {
	accdto, err := h.accRPC.GetAccountByIdRPC(ctx, nurseId.String())
	if err != nil {
		return nil, err
	}

	reldto, err := h.queryRepo.FindById(ctx, nurseId)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.NewBadRequestError().WithReason("nurse-id not found")
		}
		return nil, common.NewInternalServerError().
			WithReason("cannot get nurse info").WithInner(err.Error())
	}

	resp := &ResponseProfileDTO{
		accdto,
		toNurseDTO(reldto),
	}

	return resp, nil
}
