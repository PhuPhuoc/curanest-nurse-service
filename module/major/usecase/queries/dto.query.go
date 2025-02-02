package majorqueries

import (
	majordomain "github.com/PhuPhuoc/curanest-nurse-service/module/major/domain"
	"github.com/google/uuid"
)

type MajorDTO struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func toDTO(data *majordomain.Major) MajorDTO {
	dto := MajorDTO{
		Id:   data.GetID(),
		Name: data.GetName(),
	}
	return dto
}
