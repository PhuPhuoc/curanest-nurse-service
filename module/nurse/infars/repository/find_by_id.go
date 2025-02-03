package nurserepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
	"github.com/google/uuid"
)

func (repo *nurseRepo) FindById(ctx context.Context, id uuid.UUID) (*nursedomain.Nurse, error) {
	var dto NurseDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE, FIELD, &where)
	if err := repo.db.Get(&dto, query, id); err != nil {
		return nil, err
	}
	return dto.ToEntity()
}
