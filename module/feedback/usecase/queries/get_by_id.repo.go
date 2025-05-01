package feedbackqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type getByIdHandler struct {
	queryRepo FeedbackQueryRepo
}

func NewGetByIdHandler(queryRepo FeedbackQueryRepo) *getByIdHandler {
	return &getByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getByIdHandler) Handle(ctx context.Context, id uuid.UUID) (*FeedbackDTO, error) {
	entity, err := h.queryRepo.GetById(ctx, id)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get feedback (id: " + id.String() + ")").
			WithInner(err.Error())
	}

	return toFeedbackDTO(entity), nil
}
