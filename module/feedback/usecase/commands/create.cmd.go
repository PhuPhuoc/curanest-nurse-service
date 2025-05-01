package feedbackcommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
)

type createFeedbackHandler struct {
	cmdRepo FeedbackCommandRepo
}

func NewCreateFeedbackHandler(cmdRepo FeedbackCommandRepo) *createFeedbackHandler {
	return &createFeedbackHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createFeedbackHandler) Handle(ctx context.Context, dto *CreateFeedbackCmdDTO) error {
	newId := common.GenUUID()

	newEntity, _ := feedbackdomain.NewFeedback(
		newId,
		dto.NurseId,
		dto.MedicalRecordId,
		dto.PatientName,
		dto.Service,
		dto.Content,
		feedbackdomain.EnumFeedbackStart(dto.Star),
		nil,
	)

	if err := h.cmdRepo.Create(ctx, newEntity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create feedback").
			WithInner(err.Error())
	}
	return nil
}
