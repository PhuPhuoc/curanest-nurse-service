package nursecommands

import (
	"context"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type Commands struct {
	CreateNurse        *createNurseAccountHandler
	UpdateNurseToStaff *updateNurseToStaffHandler
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
		UpdateNurseToStaff: NewUpdateNurseToStaffHandler(
			b.BuildNurseCmdRepo(),
			b.BuildExternalAccountServiceInCmd(),
		),
	}
}

type NurseCommandRepo interface {
	Create(ctx context.Context, entity *nursedomain.Nurse) error

	UpdateNurseToStaff(ctx context.Context, nurseId uuid.UUID, categoryId uuid.UUID) error
	UpdateStaffToNurse(ctx context.Context, nurseId uuid.UUID) error
}
type ExternalAccountService interface {
	CreateAccountRPC(ctx context.Context, entity *AccountInfoDTO) (*ResponseCreateAccountDTO, error)
	HardDeleteAccountProfileRPC(ctx context.Context, accId string) error
	UpdateRoleNurseToStaffRPC(ctx context.Context, accId uuid.UUID, payload *UpdateRoleRequestRPC) error
}
