package nursehttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get detail of nursing
// @Description	get detail of nursing
// @Tags			nurse
// @Accept			json
// @Produce		json
// @Param			nurse-id	path		string					true	"Nursing ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/nurses/{nurse-id} [get]
func (s *nurseHttpService) handleGetNursingDetail() gin.HandlerFunc {
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
		res, err := s.query.GetNursingDetail.Handle(ctx, nurseUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, res)
	}
}
