package builder

import (
	"github.com/jmoiron/sqlx"

	feedbackrepository "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/infars/repository"
	feedbackcommands "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/usecase/commands"
	feedbackqueries "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/usecase/queries"
)

type builderOfFeedback struct {
	db *sqlx.DB
}

func NewFeedbackBuilder(db *sqlx.DB) builderOfFeedback {
	return builderOfFeedback{db: db}
}

func (s builderOfFeedback) BuildFeedbackCmdRepo() feedbackcommands.FeedbackCommandRepo {
	return feedbackrepository.NewFeedbackRepo(s.db)
}

func (s builderOfFeedback) BuildFeedbackQueryRepo() feedbackqueries.FeedbackQueryRepo {
	return feedbackrepository.NewFeedbackRepo(s.db)
}
