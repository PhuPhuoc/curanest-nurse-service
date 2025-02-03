package nurserepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

func (repo *nurseRepo) Create(ctx context.Context, entity *nursedomain.Nurse) error {
	accdto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, FIELD, nil)
	if _, err := repo.db.NamedExec(query, accdto); err != nil {
		return err
	}
	return nil
}
