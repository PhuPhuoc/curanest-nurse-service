package nursequeries

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

type getByIdHandler struct {
	queryRepo NurseQueryRepo
}

func NewGetByIdHandler(queryRepo NurseQueryRepo) *getByIdHandler {
	return &getByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getByIdHandler) Handle(ctx context.Context, nurseId uuid.UUID) (*NurseInfoDTO, error) {
	reldto, err := h.queryRepo.FindById(ctx, nurseId)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.NewBadRequestError().WithReason("can't find nurse with id: " + nurseId.String())
		}
		return nil, common.NewInternalServerError().
			WithReason("cannot get nurse info").
			WithInner(err.Error())
	}

	return toDTO(reldto), nil
}
