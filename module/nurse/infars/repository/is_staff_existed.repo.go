package nurserepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

func (repo *nurseRepo) IsStaffExisted(ctx context.Context, staffId uuid.UUID) (bool, error) {
	where := `staff_id=?`
	query := common.GenerateSQLQueries(common.SELECT_EXIST, STAFF_TABLE, STAFF_FIELD, &where)
	var existed bool

	if err := repo.db.GetContext(ctx, &existed, query, staffId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return existed, nil
}
