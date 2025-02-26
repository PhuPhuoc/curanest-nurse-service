package nurserepository

import (
	"context"

	"github.com/google/uuid"
)

func (repo *nurseRepo) GetNurseService(ctx context.Context, nurseId uuid.UUID) ([]uuid.UUID, error) {
	serIds := []uuid.UUID{}

	query := `select service_id from nursing_service where nursing_id=?`
	if err := repo.db.SelectContext(ctx, &serIds, query, nurseId); err != nil {
		return nil, err
	}

	return serIds, nil
}
