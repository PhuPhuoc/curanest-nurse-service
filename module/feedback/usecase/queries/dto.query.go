package feedbackqueries

import (
	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
	"github.com/google/uuid"
)

type FeedbackDTO struct {
	Id              uuid.UUID `json:"id"`
	NurseId         uuid.UUID `json:"nurse-id"`
	MedicalRecordId uuid.UUID `json:"medical-record-id"`
	PatientName     string    `json:"patient-name"`
	Service         string    `json:"service"`
	Star            string    `json:"star"`
	Content         string    `json:"content"`
}

func toFeedbackDTO(data *feedbackdomain.Feedback) *FeedbackDTO {
	dto := &FeedbackDTO{
		Id:              data.GetID(),
		NurseId:         data.GetID(),
		MedicalRecordId: data.GetMedicalRecordID(),
		PatientName:     data.GetPatientName(),
		Service:         data.GetService(),
		Star:            data.GetStar().String(),
		Content:         data.GetContent(),
	}
	return dto
}
