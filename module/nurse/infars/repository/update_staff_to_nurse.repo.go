package nurserepository

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

func (repo *nurseRepo) UpdateStaffToNurse(ctx context.Context, nurseId uuid.UUID) error {
	where := `nurse_id=?`
	query := common.GenerateSQLQueries(common.HARD_DELETE, STAFF_TABLE, STAFF_FIELD, &where)

	result, err := repo.db.ExecContext(ctx, query, nurseId)
	if err != nil {
		return err
	}

	numAffected, _ := result.RowsAffected()
	if numAffected == 0 {
		return common.ErrNoRecordsAreChanged
	}
	return nil
}
