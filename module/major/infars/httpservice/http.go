package majorhttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-nurse-service/middleware"
	majorcommands "github.com/PhuPhuoc/curanest-nurse-service/module/major/usecase/commands"
	majorqueries "github.com/PhuPhuoc/curanest-nurse-service/module/major/usecase/queries"
)

type majorHttpService struct {
	cmd   majorcommands.Commands
	query majorqueries.Queries
	auth  middleware.AuthClient
}

func NewPatientHTTPService(cmd majorcommands.Commands, query majorqueries.Queries) *majorHttpService {
	return &majorHttpService{
		cmd:   cmd,
		query: query,
	}
}

func (s *majorHttpService) AddAuth(auth middleware.AuthClient) *majorHttpService {
	s.auth = auth
	return s
}

func (s *majorHttpService) Routes(g *gin.RouterGroup) {
	major_route := g.Group("/majors")
	{
		major_route.POST(
			"",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleCreateMajor(),
		)
		major_route.GET(
			"",
			s.handleGetRoles(),
		)
	}
}
