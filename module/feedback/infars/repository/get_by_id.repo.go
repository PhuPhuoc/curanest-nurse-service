package feedbackrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
	"github.com/google/uuid"
)

func (repo *feedbackRepo) GetById(ctx context.Context, id uuid.UUID) (*feedbackdomain.Feedback, error) {
	var dto FeedbackDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE, GET_FIELD, &where)
	if err := repo.db.Get(&dto, query, id); err != nil {
		return nil, err
	}
	return dto.ToEntity()
}
