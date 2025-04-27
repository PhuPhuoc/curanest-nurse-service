package nurserepository

import (
	"context"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
	"github.com/google/uuid"
)

/*
*

	"id",
	"nurse_picture",
	"nurse_name",
	"gender",
	"current_work_place",
	"rate",

*
*/
func (repo *nurseRepo) GetByServiceId(ctx context.Context, serviceId uuid.UUID) ([]nursedomain.Nurse, error) {
	dtos := []NurseDTO{}

	query := `
	select n.id, n.nurse_picture, n.nurse_name, n.gender, n.current_work_place, n.rate
	from nursing n
	join nursing_service ns 
	on n.id = ns.nursing_id
	where ns.service_id = ?
	`
	if err := repo.db.Select(&dtos, query, serviceId); err != nil {
		return nil, err
	}

	entities := make([]nursedomain.Nurse, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}
	return entities, nil
}
