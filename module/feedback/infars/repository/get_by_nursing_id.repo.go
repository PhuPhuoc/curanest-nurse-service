package feedbackrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
	"github.com/google/uuid"
)

func (repo *feedbackRepo) GetByNursingId(ctx context.Context, nursingId uuid.UUID) ([]feedbackdomain.Feedback, error) {
	where := "nursing_id =?"
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE, GET_FIELD, &where)
	query = query + " order by created_at desc"
	var dtos []FeedbackDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, nursingId); err != nil {
		return []feedbackdomain.Feedback{}, err
	}

	entities := make([]feedbackdomain.Feedback, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}

	return entities, nil
}
