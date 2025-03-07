package nursequeries

import (
	"context"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type Queries struct {
	GetMe                   *getMyProfileHandler
	GetById                 *getByIdHandler
	GetNursingDetail        *getNursingDetailHandler
	GetNursingPrivateDetail *getNursingPrivateDetailHandler
	GetWithFilter           *getNurseWithFilterHandler

	GetStaffByIds *getStaffByIdsHandler

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

		GetNursingDetail: NewGetNursingDetailHandler(
			b.BuildNurseQueryRepo(),
		),
		GetNursingPrivateDetail: NewGetNursingPrivateDetailHandler(
			b.BuildNurseQueryRepo(),
			b.BuildExternalAccountServiceInQuery(),
		),

		GetWithFilter: NewGetNursesWithFilterHandler(
			b.BuildNurseQueryRepo(),
		),

		GetStaffByIds: NewGetStaffByIdsHandler(
			b.BuildNurseQueryRepo(),
		),

		GetNusingServiceIds: NewGetNursingServiceHandler(
			b.BuildNurseQueryRepo(),
		),
	}
}

type NurseQueryRepo interface {
	FindById(ctx context.Context, id uuid.UUID) (*nursedomain.Nurse, error)
	GetByFilter(ctx context.Context, requestQuery *NurseRequestQueryDTO) ([]nursedomain.Nurse, error)

	GetStaffByIds(ctx context.Context, ids StaffIdsQueryDTO) ([]nursedomain.Nurse, error)
	GetNurseService(ctx context.Context, nurseId uuid.UUID) ([]uuid.UUID, error)
}

type ExternalAccountService interface {
	GetAccountProfileRPC(ctx context.Context) (*ResponseAccountDTO, error)
	GetAccountByIdRPC(ctx context.Context, accId string) (*ResponseAccountDTO, error)
}
