package builder

import (
	"github.com/jmoiron/sqlx"

	majorrepository "github.com/PhuPhuoc/curanest-nurse-service/module/major/infars/repository"
	majorcommands "github.com/PhuPhuoc/curanest-nurse-service/module/major/usecase/commands"
	majorqueries "github.com/PhuPhuoc/curanest-nurse-service/module/major/usecase/queries"
)

type builderOfMajor struct {
	db *sqlx.DB
}

func NewMajorBuilder(db *sqlx.DB) builderOfMajor {
	return builderOfMajor{db: db}
}

func (s builderOfMajor) BuildMajorCmdRepo() majorcommands.MajorCommandRepo {
	return majorrepository.NewMajorRepo(s.db)
}

func (s builderOfMajor) BuildMajorQueryRepo() majorqueries.MajorQueryRepo {
	return majorrepository.NewMajorRepo(s.db)
}
