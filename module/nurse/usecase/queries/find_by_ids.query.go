package nursequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

type getStaffByIdsHandler struct {
	queryRepo NurseQueryRepo
}

func NewGetStaffByIdsHandler(queryRepo NurseQueryRepo) *getStaffByIdsHandler {
	return &getStaffByIdsHandler{
		queryRepo: queryRepo,
	}
}

func (h *getStaffByIdsHandler) Handle(ctx context.Context, ids StaffIdsQueryDTO) ([]StaffDTO, error) {
	enties, err := h.queryRepo.GetStaffByIds(ctx, ids)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list staff info").
			WithInner(err.Error())
	}
	dtos := make([]StaffDTO, len(enties))
	for i := range enties {
		dtos[i] = *toStaffDTO(&enties[i])
	}

	return dtos, nil
}
