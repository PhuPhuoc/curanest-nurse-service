package nurserpcservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/middleware"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
	"github.com/gin-gonic/gin"
)

type nursingRPCService struct {
	query nursequeries.Queries
	auth  middleware.AuthClient
}

func NewNursingRPCService(query nursequeries.Queries) *nursingRPCService {
	return &nursingRPCService{
		query: query,
	}
}

func (s *nursingRPCService) AddAuth(auth middleware.AuthClient) *nursingRPCService {
	s.auth = auth
	return s
}

func (s *nursingRPCService) Routes(g *gin.RouterGroup) {
	nursing_route := g.Group("/nurses")
	{
		nursing_route.POST(
			"/by-ids",
			middleware.RequireAuth(s.auth),
			s.handleGetStaffInfoByIds(),
		)
		nursing_route.GET(
			"/service/:service-id",
			// middleware.RequireAuth(s.auth),
			s.handleGetNuringByServiceId(),
		)
	}
}
