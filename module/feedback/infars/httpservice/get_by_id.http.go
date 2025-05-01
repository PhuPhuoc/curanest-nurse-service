package feedbackhttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get feedback
// @Description	get feedback
// @Tags			feedbacks
// @Accept			json
// @Produce		json
// @Param			feedback-id	path		string					true	"feeedback ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/feedbacks/{feedback-id} [get]
// @Security		ApiKeyAuth
func (s *feedbackHttpService) handleGetFeedbackById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var feedbackUUID uuid.UUID
		var err error
		if feedbackId := ctx.Param("feedback-id"); feedbackId != "" {
			feedbackUUID, err = uuid.Parse(feedbackId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("nursing-id invalid (not a UUID)"))
				return
			}
		}
		entity, err := s.query.GetById.Handle(ctx.Request.Context(), feedbackUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, entity)
	}
}
