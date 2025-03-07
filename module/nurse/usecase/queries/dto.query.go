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
	Id          uuid.UUID `json:"id,omitempty"`
	Role        string    `json:"role,omitempty"`
	FullName    string    `json:"full-name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone-number,omitempty"`
	CreatedAt   time.Time `json:"created-at,omitempty"`
}

type NurseDTO struct {
	NurseId          uuid.UUID `json:"nurse-id,omitempty"`
	NursePicture     string    `json:"nurse-picture,omitempty"`
	NurseName        string    `json:"nurse-name,omitempty"`
	Gender           bool      `json:"gender,omitempty"`
	Dob              string    `json:"dob,omitempty"`
	Address          string    `json:"address,omitempty"`
	Ward             string    `json:"ward,omitempty"`
	District         string    `json:"district,omitempty"`
	City             string    `json:"city,omitempty"`
	CurrentWorkPlace string    `json:"current-work-place,omitempty"`
	EducationLevel   string    `json:"education-level,omitempty"`
	Experience       string    `json:"experience,omitempty"`
	Certificate      string    `json:"certificate,omitempty"`
	Google_Drive_URL string    `json:"google-drive-url,omitempty"`
	Slogan           string    `json:"slogan,omitempty"`
	Rate             float64   `json:"rate,omitempty"`
}

func toNurseDTOForRelative(data *nursedomain.Nurse) *NurseDTO {
	dto := &NurseDTO{
		NurseId:          data.GetID(),
		NursePicture:     data.GetNursePicture(),
		NurseName:        data.GetNurseName(),
		Gender:           data.GetGender(),
		City:             data.GetCity(),
		CurrentWorkPlace: data.GetCurrentWorkPlace(),
		EducationLevel:   data.GetEducationLevel(),
		Experience:       data.GetExperience(),
		Certificate:      data.GetCertificate(),
		Slogan:           data.GetSlogan(),
		Rate:             data.GetRate(),
	}
	return dto
}

func toNurseDTO(data *nursedomain.Nurse) *NurseDTO {
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

type GetNurseDTO struct {
	NurseId          uuid.UUID `json:"nurse-id"`
	NursePicture     string    `json:"nurse-picture"`
	NurseName        string    `json:"nurse-name"`
	Gender           bool      `json:"gender"`
	CurrentWorkPlace string    `json:"current-work-place"`
	Rate             float64   `json:"rate"`
}

func toGetNurseDTO(data *nursedomain.Nurse) *GetNurseDTO {
	dto := &GetNurseDTO{
		NurseId:          data.GetID(),
		NursePicture:     data.GetNursePicture(),
		NurseName:        data.GetNurseName(),
		Gender:           data.GetGender(),
		CurrentWorkPlace: data.GetCurrentWorkPlace(),
		Rate:             data.GetRate(),
	}
	return dto
}
