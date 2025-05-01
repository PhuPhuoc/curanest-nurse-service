package feedbackqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/google/uuid"
)

type getByMedicalRecordIdHandler struct {
	queryRepo FeedbackQueryRepo
}

func NewGetByMedicalRecordIdHandler(queryRepo FeedbackQueryRepo) *getByMedicalRecordIdHandler {
	return &getByMedicalRecordIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getByMedicalRecordIdHandler) Handle(ctx context.Context, medicalRecordId uuid.UUID) (*FeedbackDTO, error) {
	entity, err := h.queryRepo.GetByMedicalRecordId(ctx, medicalRecordId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get feedback (medical-record-id: " + medicalRecordId.String() + ")").
			WithInner(err.Error())
	}

	return toFeedbackDTO(entity), nil
}
