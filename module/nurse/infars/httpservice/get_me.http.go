package nursehttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
)

//	@Summary		get nurse profile (me)
//	@Description	get nurse profile (me)
//	@Tags			nurse
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"data"
//	@Failure		400	{object}	error					"Bad request error"
//	@Router			/api/v1/nurses/me [get]
//	@Security		ApiKeyAuth
func (s *nurseHttpService) handleGetMe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dto, err := s.query.GetMe.Handle(ctx.Request.Context())
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, dto)
	}
}
