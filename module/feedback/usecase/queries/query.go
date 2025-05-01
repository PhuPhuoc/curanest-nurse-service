package feedbackqueries

type Queries struct{}

type Builder interface {
	BuildFeedbackQueryRepo() FeedbackQueryRepo
}

func NewFeedbackQueryWithBuilder(b Builder) Queries {
	return Queries{}
}

type FeedbackQueryRepo interface{}
