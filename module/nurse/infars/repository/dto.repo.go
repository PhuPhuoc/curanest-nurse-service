package nurserepository

import (
	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

var (
	TABLE = `nursing`

	FIELD = []string{
		"id",
		"nurse_picture",
		"nurse_name",
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
		"nurse_picture",
		"nurse_name",
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

	STAFF_TABLE = `staff_category`

	STAFF_FIELD = []string{
		"staff_id",
		"category_id",
	}
)

type MajorStaff struct {
	StaffId    uuid.UUID `db:"staff_id"`
	CategoryId uuid.UUID `db:"category_id"`
}

type NurseDTO struct {
	Id               uuid.UUID `db:"id"`
	Gender           bool      `db:"gender"`
	NursePicture     string    `db:"nurse_picture"`
	NurseName        string    `db:"nurse_name"`
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
		dto.Gender,
		dto.NursePicture,
		dto.NurseName,
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
		Gender:           data.GetGender(),
		NursePicture:     data.GetNursePicture(),
		NurseName:        data.GetNurseName(),
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
