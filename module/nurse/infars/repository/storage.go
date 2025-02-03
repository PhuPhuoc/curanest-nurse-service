package nurserepository

import "github.com/jmoiron/sqlx"

type nurseRepo struct {
	db *sqlx.DB
}

func NewNurseRepo(db *sqlx.DB) *nurseRepo {
	return &nurseRepo{
		db: db,
	}
}
