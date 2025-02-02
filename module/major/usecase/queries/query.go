package majorqueries

import (
	"context"

	majordomain "github.com/PhuPhuoc/curanest-nurse-service/module/major/domain"
)

type Queries struct {
	GetAllMajors *getMajorsHandler
}

type Builder interface {
	BuildMajorQueryRepo() MajorQueryRepo
}

func NewMajorQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetAllMajors: NewGetMajorHandler(
			b.BuildMajorQueryRepo(),
		),
	}
}

type MajorQueryRepo interface {
	GetAllMajors(ctx context.Context) ([]majordomain.Major, error)
}
