package majorrepository

import "github.com/jmoiron/sqlx"

type majorRepo struct {
	db *sqlx.DB
}

func NewMajorRepo(db *sqlx.DB) *majorRepo {
	return &majorRepo{
		db: db,
	}
}
