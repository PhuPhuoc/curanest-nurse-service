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

	majorid, err := uuid.Parse(dto.MajorId)
	if err != nil {
		return common.NewBadRequestError().WithReason("your major-id is invalid - it must be a uuid")
	}

	// 2. create record in table relatives
	accid := uuid.MustParse(resp.Id)
	entity, _ := nursedomain.NewNurse(
		accid,
		majorid,
		dto.Gender,
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
		"",
	)
	if err = h.cmdRepo.Create(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create relatives info").
			WithInner(err.Error())
	}

	return nil
}
