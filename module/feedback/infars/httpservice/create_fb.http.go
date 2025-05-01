package feedbackhttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	feedbackcommands "github.com/PhuPhuoc/curanest-nurse-service/module/feedback/usecase/commands"
	"github.com/gin-gonic/gin"
)

//	@Summary		create feedback for nursing
//	@Description	create feedback for nursing
//	@Tags			feedbacks
//	@Accept			json
//	@Produce		json
//	@Param			create	form		body					feedbackcommands.CreateFeedbackCmdDTO	true	"feedback creation data"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/api/v1/feedbacks [post]
//	@Security		ApiKeyAuth
func (s *feedbackHttpService) handleCreateFeedback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto feedbackcommands.CreateFeedbackCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreateFeedback.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
