package nursequeries

import (
	"context"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type Queries struct {
	GetMe *getMyProfileHandler
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
	}
}

type NurseQueryRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*nursedomain.Nurse, error)
}

type ExternalAccountService interface {
	GetAccountProfileRPC(ctx context.Context) (*ResponseAccountDTO, error)
}
