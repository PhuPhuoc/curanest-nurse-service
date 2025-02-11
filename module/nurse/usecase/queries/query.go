package nursequeries

import (
	"context"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type Queries struct {
	GetMe        *getMyProfileHandler
	GetById      *getByIdHandler
	IsNurseStaff *isNurseStaffHandler
}

type Builder interface {
	BuildNurseQueryRepo() NurseQueryRepo
	BuildExternalAccountServiceInQuery() ExternalAccountService
}

func NewNurseQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetMe: NewGetMyNurseAccountHandler(
			b.BuildNurseQueryRepo(),
			b.BuildExternalAccountServiceInQuery(),
		),
		GetById: NewGetByIdHandler(
			b.BuildNurseQueryRepo(),
		),
		IsNurseStaff: NewIsNurseStaffHandler(
			b.BuildNurseQueryRepo(),
		),
	}
}

type NurseQueryRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*nursedomain.Nurse, error)
	IsStaffExisted(ctx context.Context, staffId uuid.UUID) (bool, error)
}

type ExternalAccountService interface {
	GetAccountProfileRPC(ctx context.Context) (*ResponseAccountDTO, error)
}
