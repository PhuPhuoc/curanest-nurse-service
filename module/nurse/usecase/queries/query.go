package nursequeries

import (
	"context"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type Queries struct {
	GetMe   *getMyProfileHandler
	GetById *getByIdHandler

	GetNusingServiceIds *getNursingServiceHandler
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

		GetNusingServiceIds: NewGetNursingServiceHandler(
			b.BuildNurseQueryRepo(),
		),
	}
}

type NurseQueryRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*nursedomain.Nurse, error)

	GetNurseService(ctx context.Context, nurseId uuid.UUID) ([]uuid.UUID, error)
}

type ExternalAccountService interface {
	GetAccountProfileRPC(ctx context.Context) (*ResponseAccountDTO, error)
}
