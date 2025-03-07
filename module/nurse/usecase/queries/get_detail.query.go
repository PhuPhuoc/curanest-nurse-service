package nursequeries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type getNursingDetailHandler struct {
	queryRepo NurseQueryRepo
}

func NewGetNursingDetailHandler(queryRepo NurseQueryRepo) *getNursingDetailHandler {
	return &getNursingDetailHandler{
		queryRepo: queryRepo,
	}
}

func (h *getNursingDetailHandler) Handle(ctx context.Context, nurseId uuid.UUID) (*NurseDTO, error) {
	reldto, err := h.queryRepo.FindById(ctx, nurseId)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.NewBadRequestError().WithReason("nurse-id not found")
		}
		return nil, common.NewInternalServerError().
			WithReason("cannot get nurse info").WithInner(err.Error())
	}

	resp := toNurseDTOForRelative(reldto)

	return resp, nil
}
