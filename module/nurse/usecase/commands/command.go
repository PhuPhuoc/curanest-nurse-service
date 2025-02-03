package nursecommands

import (
	"context"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type Commands struct {
	CreateNurse *createNurseAccountHandler
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
	}
}

type NurseCommandRepo interface {
	Create(ctx context.Context, entity *nursedomain.Nurse) error
}
type ExternalAccountService interface {
	CreateAccountRPC(ctx context.Context, entity *AccountInfoDTO) (*ResponseCreateAccountDTO, error)
}
