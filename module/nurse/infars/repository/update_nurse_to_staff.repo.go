package nurserepository

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

func (repo *nurseRepo) UpdateNurseToStaff(ctx context.Context, nurseId uuid.UUID, categoryId uuid.UUID) error {
	query := common.GenerateSQLQueries(common.INSERT, STAFF_TABLE, STAFF_FIELD, nil)
	dto := &MajorStaff{
		nurseId,
		categoryId,
	}

	if _, err := repo.db.NamedExecContext(ctx, query, dto); err != nil {
		return err
	}
	return nil
}
