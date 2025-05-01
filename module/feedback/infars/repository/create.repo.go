package feedbackrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
)

func (repo *feedbackRepo) Create(ctx context.Context, entity *feedbackdomain.Feedback) error {
	accdto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, CREATE_FIELD, nil)
	if _, err := repo.db.NamedExec(query, accdto); err != nil {
		return err
	}
	return nil
}
