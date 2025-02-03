package nursequeries

import (
	"time"

	"github.com/google/uuid"

	nursedomain "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/domain"
)

type ResponseProfileDTO struct {
	*ResponseAccountDTO
	*NurseInfoDTO
}

type ResponseAccountDTO struct {
	Id          uuid.UUID `json:"id"`
	Role        string    `json:"role"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created-at"`
}
type NurseInfoDTO struct {
	Gender           bool   `json:"gender"`
	Dob              string `json:"dob"`
	Address          string `json:"address"`
	Ward             string `json:"ward"`
	District         string `json:"district"`
	City             string `json:"city"`
	CurrentWorkPlace string `json:"current-work-place"`
	EducationLevel   string `json:"education-level"`
	Experience       string `json:"experience"`
	Certificate      string `json:"certificate"`
	Google_Drive_URL string `json:"google-drive-url"`
	Slogan           string `json:"slogan"`
	Rate             string `json:"rate"`
}

func toDTO(data *nursedomain.Nurse) *NurseInfoDTO {
	dto := &NurseInfoDTO{
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
