package majorrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	majordomain "github.com/PhuPhuoc/curanest-nurse-service/module/major/domain"
)

func (repo *majorRepo) Create(ctx context.Context, entity *majordomain.Major) error {
	accdto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, FIELD, nil)
	if _, err := repo.db.NamedExec(query, accdto); err != nil {
		return err
	}
	return nil
}
