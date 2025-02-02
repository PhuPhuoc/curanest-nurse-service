package majorhttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	majorcommands "github.com/PhuPhuoc/curanest-nurse-service/module/major/usecase/commands"
)

// @Summary		create new major - admin
// @Description	create new major
// @Tags			major
// @Accept			json
// @Produce		json
// @Param			create	form		body					majorcommands.CreateMajorCmdDTO	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/majors [post]
// @Security		ApiKeyAuth
func (s *majorHttpService) handleCreateMajor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto majorcommands.CreateMajorCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.cmd.CreateMajor.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseCreated(ctx)
	}
}
