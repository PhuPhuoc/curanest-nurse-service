package feedbackrepository

import "github.com/jmoiron/sqlx"

type feedbackRepo struct {
	db *sqlx.DB
}

func NewFeedbackRepo(db *sqlx.DB) *feedbackRepo {
	return &feedbackRepo{
		db: db,
	}
}
