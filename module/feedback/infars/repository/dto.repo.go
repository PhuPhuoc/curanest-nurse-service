package feedbackrepository

import (
	"time"

	feedbackdomain "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/domain"
	"github.com/google/uuid"
)

var (
	TABLE = `feedbacks`

	CREATE_FIELD = []string{
		"id",
		"nursing_id",
		"medical_record_id",
		"patient_name",
		"service",
		"star",
		"content",
	}

	GET_FIELD = []string{
		"id",
		"nursing_id",
		"medical_record_id",
		"patient_name",
		"service",
		"star",
		"content",
	}

	UPDATE_FIELD = []string{
		"service",
		"star",
		"content",
	}
)

type FeedbackDTO struct {
	Id              uuid.UUID  `db:"id"`
	NursingId       uuid.UUID  `db:"nursing_id"`
	MedicalRecordId uuid.UUID  `db:"medical_record_id"`
	PatientName     string     `db:"patient_name"`
	Service         string     `db:"service"`
	Star            string     `db:"star"`
	Content         string     `db:"content"`
	CreatedAt       *time.Time `db:"created_at"`
}

func (dto *FeedbackDTO) ToEntity() (*feedbackdomain.Feedback, error) {
	return feedbackdomain.NewFeedback(
		dto.Id,
		dto.NursingId,
		dto.MedicalRecordId,
		dto.PatientName,
		dto.Service,
		dto.Content,
		feedbackdomain.EnumFeedbackStart(dto.Star),
		dto.CreatedAt,
	)
}

func ToDTO(data *feedbackdomain.Feedback) *FeedbackDTO {
	return &FeedbackDTO{
		Id:              data.GetID(),
		NursingId:       data.GetNursingID(),
		MedicalRecordId: data.GetMedicalRecordID(),
		PatientName:     data.GetPatientName(),
		Service:         data.GetService(),
		Star:            data.GetStar().String(),
	}
}
