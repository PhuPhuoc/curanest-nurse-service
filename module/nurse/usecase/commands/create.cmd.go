package nursecommands

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type createNurseAccountHandler struct {
	cmdRepo    NurseCommandRepo
	accService ExternalAccountService
}

func NewCreateNurseAccountHandler(cmdRepo NurseCommandRepo, accService ExternalAccountService) *createNurseAccountHandler {
	return &createNurseAccountHandler{
		cmdRepo:    cmdRepo,
		accService: accService,
	}
}

func (h *createNurseAccountHandler) Handle(ctx context.Context, dto *CreateNurseAccountCmdDTO) error {
	// 1. call external service
	accdto := &AccountInfoDTO{
		RoleName:    "nurse",
		FullName:    dto.FullName,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
		Password:    dto.Password,
	}

	resp, err := h.accService.CreateAccountRPC(ctx, accdto)
	if err != nil {
		return err
	}

	if resp.Id == "" {
		return common.NewInternalServerError().
			WithReason("cannot create account for relatives - cannot get account id")
	}
	accid := uuid.MustParse(resp.Id)

	// 2. create record in table relatives
	entity, _ := nursedomain.NewNurse(
		accid,
		dto.Gender,
		dto.NursePicture,
		dto.FullName,
		dto.CitizenId,
		dto.Dob,
		dto.Address,
		dto.Ward,
		dto.District,
		dto.City,
		dto.CurrentWorkPlace,
		dto.EducationLevel,
		dto.Experience,
		dto.Certificate,
		dto.GoogleDriveUrl,
		dto.Slogan,
		4.9777,
	)
	if err = h.cmdRepo.Create(ctx, entity); err != nil {
		_ = h.accService.HardDeleteAccountProfileRPC(ctx, accid.String())
		return common.NewInternalServerError().
			WithReason("cannot create relatives info").
			WithInner(err.Error())
	}

	return nil
}
