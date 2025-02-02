package majorhttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
)

// @Summary		get major
// @Description	get major
// @Tags			major
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/api/v1/majors [get]
func (s *majorHttpService) handleGetRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := s.query.GetAllMajors.Handle(c.Request.Context())
		if err != nil {
			common.ResponseError(c, err)
			return
		}

		common.ResponseSuccess(c, result)
	}
}
