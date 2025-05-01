package feedbackdomain

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	id              uuid.UUID
	nursingId       uuid.UUID
	medicalRecordId uuid.UUID
	patientName     string
	service         string
	star            FeedbackStart
	content         string
	createdAt       *time.Time
}

func NewFeedback(
	id, nursingId, medicalRecordId uuid.UUID,
	patientName, service, content string,
	star FeedbackStart,
	createdAt *time.Time,
) (*Feedback, error) {
	return &Feedback{
		id:              id,
		nursingId:       nursingId,
		medicalRecordId: medicalRecordId,
		patientName:     patientName,
		service:         service,
		star:            star,
		content:         content,
		createdAt:       createdAt,
	}, nil
}

func (f *Feedback) GetID() uuid.UUID {
	return f.id
}

func (f *Feedback) GetNursingID() uuid.UUID {
	return f.nursingId
}

func (f *Feedback) GetMedicalRecordID() uuid.UUID {
	return f.medicalRecordId
}

func (f *Feedback) GetPatientName() string {
	return f.patientName
}

func (f *Feedback) GetService() string {
	return f.service
}

func (f *Feedback) GetStar() FeedbackStart {
	return f.star
}

func (f *Feedback) GetContent() string {
	return f.content
}

func (f *Feedback) GetCreatedAt() *time.Time {
	return f.createdAt
}

type FeedbackStart int

const (
	FeedbackStarOne   FeedbackStart = 1
	FeedbackStarTwo   FeedbackStart = 2
	FeedbackStarThree FeedbackStart = 3
	FeedbackStarFour  FeedbackStart = 4
	FeedbackStarFive  FeedbackStart = 5
)

// String converts FeedbackStart to string representation
func (f FeedbackStart) String() string {
	switch f {
	case FeedbackStarOne:
		return "1"
	case FeedbackStarTwo:
		return "2"
	case FeedbackStarThree:
		return "3"
	case FeedbackStarFour:
		return "4"
	case FeedbackStarFive:
		return "5"
	default:
		return "unknown"
	}
}

// EnumFeedbackStart converts a string or number to FeedbackStart
func EnumFeedbackStart(s string) FeedbackStart {
	switch s {
	case "1":
		return FeedbackStarOne
	case "2":
		return FeedbackStarTwo
	case "3":
		return FeedbackStarThree
	case "4":
		return FeedbackStarFour
	case "5":
		return FeedbackStarFive
	default:
		return FeedbackStarOne // Default to lowest rating if unknown
	}
}
