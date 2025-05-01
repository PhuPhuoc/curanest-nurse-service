package feedbackqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type getByNursingIdHandler struct {
	queryRepo FeedbackQueryRepo
}

func NewGetByNursingIdHandler(queryRepo FeedbackQueryRepo) *getByNursingIdHandler {
	return &getByNursingIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getByNursingIdHandler) Handle(ctx context.Context, nursingId uuid.UUID) ([]FeedbackDTO, error) {
	enties, err := h.queryRepo.GetByNursingId(ctx, nursingId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list feedbacks of this nursing").
			WithInner(err.Error())
	}
	dtos := make([]FeedbackDTO, len(enties))
	for i := range enties {
		dtos[i] = *toFeedbackDTO(&enties[i])
	}

	return dtos, nil
}
