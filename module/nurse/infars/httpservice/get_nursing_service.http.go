package nursehttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get service of nursing
// @Description	get service of nursing
// @Tags			nurse
// @Accept			json
// @Produce		json
// @Param			nurse-id	path		string					true	"Nursing ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/nurses/{nurse-id}/services [get]
// @Security		ApiKeyAuth
func (s *nurseHttpService) handleGetNursingServices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nurseId := ctx.Param("nurse-id")
		if nurseId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing nurse-id"))
			return
		}
		nurseUUID, err := uuid.Parse(nurseId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("nurse-id invalid (not a uuid)"))
			return
		}
		res, err := s.query.GetNusingServiceIds.Handle(ctx, nurseUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, res)
	}
}
