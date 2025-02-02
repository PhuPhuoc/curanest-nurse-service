package majorrepository

import (
	"github.com/google/uuid"

	majordomain "github.com/PhuPhuoc/curanest-nurse-service/module/major/domain"
)

var (
	TABLE        = `majors`
	FIELD        = []string{"id", "name"}
	GET_FIELD    = []string{"id", "name"}
	UPDATE_FIELD = []string{"name"}
)

type MajorDTO struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (dto *MajorDTO) ToEntity() (*majordomain.Major, error) {
	return majordomain.NewMajor(
		dto.Id,
		dto.Name,
	)
}

func ToDTO(data *majordomain.Major) *MajorDTO {
	dto := &MajorDTO{
		Id:   data.GetID(),
		Name: data.GetName(),
	}
	return dto
}
