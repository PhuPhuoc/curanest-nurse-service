package nursehttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

//	@Summary		update nurse to staff of major
//	@Description	update nurse to staff of major
//	@Tags			nurse
//	@Accept			json
//	@Produce		json
//	@Param			nurse-id	path		string					true	"Account ID (UUID)"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/nurses/{nurse-id}/staff [patch]
//	@Security		ApiKeyAuth
func (s *nurseHttpService) handleUpdateNurseToStaff() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nurseId := ctx.Param("nurse-id")
		if nurseId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing patient-id"))
			return
		}

		nurseUUID, err := uuid.Parse(nurseId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("id invalid (not a uuid)"))
			return
		}

		if err := s.query.IsNurseStaff.Handle(ctx, nurseUUID); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		nursedto, err := s.query.GetById.Handle(ctx.Request.Context(), nurseUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if err := s.command.UpdateNurseToStaff.Handle(ctx.Request.Context(), nurseUUID, nursedto.MajorId); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
