package nursequeries

import "context"

type getNurseWithFilterHandler struct {
	queryRepo NurseQueryRepo
}

func NewGetNursesWithFilterHandler(queryRepo NurseQueryRepo) *getNurseWithFilterHandler {
	return &getNurseWithFilterHandler{
		queryRepo: queryRepo,
	}
}

func (h *getNurseWithFilterHandler) Handle(ctx context.Context) error {
	return nil
}
