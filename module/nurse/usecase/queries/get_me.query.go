package nursequeries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

type getMyProfileHandler struct {
	queryRepo NurseQueryRepo
	accRPC    ExternalAccountService
}

func NewGetMyNurseAccountHandler(queryRepo NurseQueryRepo, accRPC ExternalAccountService) *getMyProfileHandler {
	return &getMyProfileHandler{
		queryRepo: queryRepo,
		accRPC:    accRPC,
	}
}

func (h *getMyProfileHandler) Handle(ctx context.Context) (*ResponseProfileDTO, error) {
	accdto, err := h.accRPC.GetAccountProfileRPC(ctx)
	if err != nil {
		return nil, err
	}

	reldto, err := h.queryRepo.FindById(ctx, accdto.Id)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.NewBadRequestError().WithReason("nurse-id not found")
		}
		return nil, common.NewInternalServerError().
			WithReason("cannot get nurse info").WithInner(err.Error())
	}

	resp := &ResponseProfileDTO{
		accdto,
		toDTO(reldto),
	}

	return resp, nil
}
