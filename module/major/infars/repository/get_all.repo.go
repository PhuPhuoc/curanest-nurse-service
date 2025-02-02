package majorrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	majordomain "github.com/PhuPhuoc/curanest-nurse-service/module/major/domain"
)

func (repo *majorRepo) GetAllMajors(ctx context.Context) ([]majordomain.Major, error) {
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE, FIELD, nil)

	var dtos []MajorDTO
	if err := repo.db.SelectContext(ctx, &dtos, query); err != nil {
		return nil, err
	}

	entities := make([]majordomain.Major, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}
	return entities, nil
}
