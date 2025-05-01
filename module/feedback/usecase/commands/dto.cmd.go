package feedbackcommands

import "github.com/google/uuid"

type CreateFeedbackCmdDTO struct {
	NurseId         uuid.UUID `json:"nurse-id" binding:"required"`
	MedicalRecordId uuid.UUID `json:"medical-record-id" binding:"required"`
	PatientName     string    `json:"patient-name" binding:"required"`
	Service         string    `json:"service" binding:"required"`
	Star            string    `json:"star" binding:"required,oneof=1 2 3 4 5"`
	Content         string    `json:"content" binding:"required"`
}
