package majorcommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	majordomain "github.com/PhuPhuoc/curanest-nurse-service/module/major/domain"
)

type createMajorHandler struct {
	cmdRepo MajorCommandRepo
}

func NewCreateMajorHandler(cmdRepo MajorCommandRepo) *createMajorHandler {
	return &createMajorHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createMajorHandler) Handle(ctx context.Context, dto *CreateMajorCmdDTO) error {
	id := common.GenUUID()
	domain, _ := majordomain.NewMajor(id, dto.Name)
	if err := h.cmdRepo.Create(ctx, domain); err != nil {
		return common.NewInternalServerError().WithMessage("cannot create major").WithInner(err.Error())
	}
	return nil
}
