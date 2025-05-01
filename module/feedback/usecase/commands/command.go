package feedbackcommands

import (
	"context"

	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
)

type Commands struct {
	CreateFeedback *createFeedbackHandler
}

type Builder interface {
	BuildFeedbackCmdRepo() FeedbackCommandRepo
}

func NewFeedbackCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateFeedback: NewCreateFeedbackHandler(
			b.BuildFeedbackCmdRepo(),
		),
	}
}

type FeedbackCommandRepo interface {
	Create(ctx context.Context, entity *feedbackdomain.Feedback) error
}
