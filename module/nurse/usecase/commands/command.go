package nursecommands

import (
	"context"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type Commands struct {
	CreateNurse *createNurseAccountHandler

	MapNurseService *mapNurseServiceHandler
}

type Builder interface {
	BuildNurseCmdRepo() NurseCommandRepo
	BuildExternalAccountServiceInCmd() ExternalAccountService
}

func NewNurseCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateNurse: NewCreateNurseAccountHandler(
			b.BuildNurseCmdRepo(),
			b.BuildExternalAccountServiceInCmd(),
		),

		MapNurseService: NewMapNurseServiceHandler(
			b.BuildNurseCmdRepo(),
		),
	}
}

type NurseCommandRepo interface {
	Create(ctx context.Context, entity *nursedomain.Nurse) error

	// map nursing with services
	MapNurseService(ctx context.Context, nurseId uuid.UUID, serviceIds []uuid.UUID) error
}
type ExternalAccountService interface {
	CreateAccountRPC(ctx context.Context, entity *AccountInfoDTO) (*ResponseCreateAccountDTO, error)
	HardDeleteAccountProfileRPC(ctx context.Context, accId string) error
}
