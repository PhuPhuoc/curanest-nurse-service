package majorcommands

import (
	"context"

	majordomain "github.com/PhuPhuoc/curanest-nurse-service/module/major/domain"
)

type Commands struct {
	CreateMajor *createMajorHandler
}

type Builder interface {
	BuildMajorCmdRepo() MajorCommandRepo
}

func NewMajorCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateMajor: NewCreateMajorHandler(b.BuildMajorCmdRepo()),
	}
}

type MajorCommandRepo interface {
	Create(ctx context.Context, entity *majordomain.Major) error
}
