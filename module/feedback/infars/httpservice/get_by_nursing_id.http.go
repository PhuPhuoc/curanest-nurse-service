package feedbackhttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		get feedback of nursing
//	@Description	get feedback of nursing
//	@Tags			feedbacks
//	@Accept			json
//	@Produce		json
//	@Param			nursing-id	path		string					true	"nursing ID (UUID)"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/feedbacks/{nursing-id} [get]
//	@Security		ApiKeyAuth
func (s *feedbackHttpService) handleGetFeedbackByNursingId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var nursingUUID uuid.UUID
		var err error
		if nursingId := ctx.Param("nursing-id"); nursingId != "" {
			nursingUUID, err = uuid.Parse(nursingId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("nursing-id invalid (not a UUID)"))
				return
			}
		}
		entites, err := s.query.GetByNursingId.Handle(ctx.Request.Context(), nursingUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, entites)
	}
}
