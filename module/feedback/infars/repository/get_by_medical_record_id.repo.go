package feedbackrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
	"github.com/google/uuid"
)

func (repo *feedbackRepo) GetByMedicalRecordId(ctx context.Context, id uuid.UUID) (*feedbackdomain.Feedback, error) {
	var dto FeedbackDTO
	where := "medical_record_id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE, GET_FIELD, &where)
	if err := repo.db.Get(&dto, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return dto.ToEntity()
}
