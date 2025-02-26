package nursehttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-nurse-service/middleware"
	nursecommands "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/commands"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
)

type nurseHttpService struct {
	command nursecommands.Commands
	query   nursequeries.Queries
	auth    middleware.AuthClient
}

func NewPatientHTTPService(command nursecommands.Commands, query nursequeries.Queries) *nurseHttpService {
	return &nurseHttpService{
		command: command,
		query:   query,
	}
}

func (s *nurseHttpService) AddAuth(auth middleware.AuthClient) *nurseHttpService {
	s.auth = auth
	return s
}

func (s *nurseHttpService) Routes(g *gin.RouterGroup) {
	nurse_route := g.Group("/nurses")
	{
		nurse_route.POST(
			"",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleCreateNurseAccount(),
		)
		nurse_route.GET(
			"/me",
			middleware.RequireAuth(s.auth),
			s.handleGetMe(),
		)
		nurse_route.POST(
			"/:nurse-id/services",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleMapNursingServices(),
		)
		nurse_route.GET(
			"/:nurse-id/services",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleGetNursingServices(),
		)
	}
}
