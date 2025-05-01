package feedbackqueries

import (
	"context"

	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
	"github.com/google/uuid"
)

type Queries struct {
	GetByNursingId *getByNursingIdHandler
}

type Builder interface {
	BuildFeedbackQueryRepo() FeedbackQueryRepo
}

func NewFeedbackQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetByNursingId: NewGetByNursingIdHandler(
			b.BuildFeedbackQueryRepo(),
		),
	}
}

type FeedbackQueryRepo interface {
	GetById(ctx context.Context, id uuid.UUID) (*feedbackdomain.Feedback, error)
	GetByNursingId(ctx context.Context, nursingId uuid.UUID) ([]feedbackdomain.Feedback, error)
}
