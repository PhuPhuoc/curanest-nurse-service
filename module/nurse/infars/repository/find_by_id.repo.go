package nurserepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

func (repo *nurseRepo) FindById(ctx context.Context, id uuid.UUID) (*nursedomain.Nurse, error) {
	var dto NurseDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE, FIELD, &where)
	if err := repo.db.Get(&dto, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}

	return dto.ToEntity()
}
