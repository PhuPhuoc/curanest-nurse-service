package nursehttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursecommands "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/commands"
	"github.com/gin-gonic/gin"
)

// @Summary		create nurse account
// @Description	create nurse account
// @Tags			nurse
// @Accept			json
// @Produce		json
// @Param			create	form		body					nursecommands.CreateNurseAccountCmdDTO	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/nurses [post]
// @Security		ApiKeyAuth
func (s *nurseHttpService) handleCreateNurseAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto nursecommands.CreateNurseAccountCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreateNurse.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
