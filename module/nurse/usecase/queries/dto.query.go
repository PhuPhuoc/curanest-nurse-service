package nursequeries

import (
	"time"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type ResponseProfileDTO struct {
	*ResponseAccountDTO
	*NurseDTO
}

type ResponseAccountDTO struct {
	Id          uuid.UUID `json:"id"`
	Role        string    `json:"role"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	CreatedAt   time.Time `json:"created-at"`
}

type NurseDTO struct {
	NurseId          uuid.UUID `json:"nurse-id"`
	NursePicture     string    `json:"nurse-picture"`
	NurseName        string    `json:"nurse-name"`
	Gender           bool      `json:"gender"`
	Dob              string    `json:"dob"`
	Address          string    `json:"address"`
	Ward             string    `json:"ward"`
	District         string    `json:"district"`
	City             string    `json:"city"`
	CurrentWorkPlace string    `json:"current-work-place"`
	EducationLevel   string    `json:"education-level"`
	Experience       string    `json:"experience"`
	Certificate      string    `json:"certificate"`
	Google_Drive_URL string    `json:"google-drive-url"`
	Slogan           string    `json:"slogan"`
	Rate             float64   `json:"rate"`
}

func toDTO(data *nursedomain.Nurse) *NurseDTO {
	dto := &NurseDTO{
		NurseId:          data.GetID(),
		NursePicture:     data.GetNursePicture(),
		NurseName:        data.GetNurseName(),
		Gender:           data.GetGender(),
		Dob:              data.GetDOB(),
		Address:          data.GetAddress(),
		Ward:             data.GetWard(),
		District:         data.GetDistrict(),
		City:             data.GetCity(),
		CurrentWorkPlace: data.GetCurrentWorkPlace(),
		EducationLevel:   data.GetEducationLevel(),
		Experience:       data.GetExperience(),
		Certificate:      data.GetCertificate(),
		Google_Drive_URL: data.GetGoogleDriveURL(),
		Slogan:           data.GetSlogan(),
		Rate:             data.GetRate(),
	}
	return dto
}

type UpdateNurseDTO struct {
	NursePicture     string  `json:"nurse-picture"`
	NurseName        string  `json:"nurse-name"`
	Gender           bool    `json:"gender"`
	CurrentWorkPlace string  `json:"current-work-place"`
	Slogan           string  `json:"slogan"`
	Rate             float64 `json:"rate"`
}

func toUpdateNurseDTO(data *nursedomain.Nurse) *UpdateNurseDTO {
	dto := &UpdateNurseDTO{
		NursePicture:     data.GetNursePicture(),
		NurseName:        data.GetNurseName(),
		Gender:           data.GetGender(),
		CurrentWorkPlace: data.GetCurrentWorkPlace(),
		Slogan:           data.GetSlogan(),
		Rate:             data.GetRate(),
	}
	return dto
}

type NursingServiceDTO struct {
	ServiceIds []uuid.UUID `json:"service-ids"`
}

type StaffIdsQueryDTO struct {
	Ids []uuid.UUID `json:"ids"`
}

type StaffDTO struct {
	NurseId      uuid.UUID `json:"nurse-id"`
	NursePicture string    `json:"nurse-picture"`
	NurseName    string    `json:"nurse-name"`
}

func toStaffDTO(data *nursedomain.Nurse) *StaffDTO {
	dto := &StaffDTO{
		NurseId:      data.GetID(),
		NursePicture: data.GetNursePicture(),
		NurseName:    data.GetNurseName(),
	}
	return dto
}
