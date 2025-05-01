package feedbackhttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/middleware"
	feedbackcommands "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/usecase/commands"
	feedbackqueries "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/usecase/queries"
	"github.com/gin-gonic/gin"
)

type feedbackHttpService struct {
	command feedbackcommands.Commands
	query   feedbackqueries.Queries
	auth    middleware.AuthClient
}

func NewfeedbackHTTPService(command feedbackcommands.Commands, query feedbackqueries.Queries) *feedbackHttpService {
	return &feedbackHttpService{
		command: command,
		query:   query,
	}
}

func (s *feedbackHttpService) AddAuth(auth middleware.AuthClient) *feedbackHttpService {
	s.auth = auth
	return s
}

func (s *feedbackHttpService) Routes(g *gin.RouterGroup) {
	feedback_route := g.Group("/feedbacks")
	{
		feedback_route.POST(
			"",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("admin"),
			s.handleCreateFeedback(),
		)
		feedback_route.GET(
			"/nursing/:nursing-id",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("admin"),
			s.handleGetFeedbackByNursingId(),
		)
		feedback_route.GET(
			"/:medical-record-id",
			// middleware.RequireAuth(s.auth),
			// middleware.RequireRole("admin"),
			s.handleGetFeedbackByMedicalRecordId(),
		)
	}
}
