package nurserepository

import (
	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

var (
	TABLE = `nurses`

	FIELD = []string{
		"id",
		"major_id",
		"gender",
		"citizen_id",
		"dob",
		"address",
		"ward",
		"district",
		"city",
		"current_work_place",
		"education_level",
		"experience",
		"certificate",
		"google_drive_url",
		"slogan",
		"rate",
	}

	UPDATE_FIELD = []string{
		"major_id",
		"gender",
		"citizen_id",
		"dob",
		"address",
		"ward",
		"district",
		"city",
		"current_work_place",
		"education_level",
		"experience",
		"certificate",
		"google_drive_url",
		"slogan",
		"rate",
	}
)

type NurseDTO struct {
	Id               uuid.UUID `db:"id"`
	MajorID          uuid.UUID `db:"major_id"`
	Gender           bool      `db:"gender"`
	CitizenID        string    `db:"citizen_id"`
	Dob              string    `db:"dob"`
	Address          string    `db:"address"`
	Ward             string    `db:"ward"`
	District         string    `db:"district"`
	City             string    `db:"city"`
	CurrentWorkPlace string    `db:"current_work_place"`
	EducationLevel   string    `db:"education_level"`
	Experience       string    `db:"experience"`
	Certificate      string    `db:"certificate"`
	GoogleDriveURL   string    `db:"google_drive_url"`
	Slogan           string    `db:"slogan"`
	Rate             string    `db:"rate"`
}

func (dto *NurseDTO) ToEntity() (*nursedomain.Nurse, error) {
	return nursedomain.NewNurse(
		dto.Id,
		dto.MajorID,
		dto.Gender,
		dto.CitizenID,
		dto.Dob,
		dto.Address,
		dto.Ward,
		dto.District,
		dto.City,
		dto.CurrentWorkPlace,
		dto.EducationLevel,
		dto.Experience,
		dto.Certificate,
		dto.GoogleDriveURL,
		dto.Slogan,
		dto.Rate,
	)
}

func ToDTO(data *nursedomain.Nurse) *NurseDTO {
	return &NurseDTO{
		Id:               data.GetID(),
		MajorID:          data.GetMajorID(),
		Gender:           data.GetGender(),
		CitizenID:        data.GetCitizenID(),
		Dob:              data.GetDOB(),
		Address:          data.GetAddress(),
		Ward:             data.GetWard(),
		District:         data.GetDistrict(),
		City:             data.GetCity(),
		CurrentWorkPlace: data.GetCurrentWorkPlace(),
		EducationLevel:   data.GetEducationLevel(),
		Experience:       data.GetExperience(),
		Certificate:      data.GetCertificate(),
		GoogleDriveURL:   data.GetGoogleDriveURL(),
		Slogan:           data.GetSlogan(),
		Rate:             data.GetRate(),
	}
}
