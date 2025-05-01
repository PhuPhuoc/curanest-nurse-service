package feedbackhttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		get feedback
//	@Description	get feedback
//	@Tags			feedbacks
//	@Accept			json
//	@Produce		json
//	@Param			medical-record-id	path		string					true	"medical-record ID (UUID)"
//	@Success		200					{object}	map[string]interface{}	"data"
//	@Failure		400					{object}	error					"Bad request error"
//	@Router			/api/v1/feedbacks/{medical-record-id} [get]
//	@Security		ApiKeyAuth
func (s *feedbackHttpService) handleGetFeedbackByMedicalRecordId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var recordUUID uuid.UUID
		var err error
		if recordId := ctx.Param("medical-record-id"); recordId != "" {
			recordUUID, err = uuid.Parse(recordId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("nursing-id invalid (not a UUID)"))
				return
			}
		}
		entity, err := s.query.GetById.Handle(ctx.Request.Context(), recordUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, entity)
	}
}
