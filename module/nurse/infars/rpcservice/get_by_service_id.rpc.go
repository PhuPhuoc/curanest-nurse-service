package nurserpcservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get staff info by ids
// @Description	get staff info by ids
// @Tags			rpc: nurse
// @Accept			json
// @Produce		json
// @Param			service-id	path		string					true	"service ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/external/rpc/nurses/service/{service-id} [get]
// @Security		ApiKeyAuth
func (s *nursingRPCService) handleGetNuringByServiceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		serviceId := ctx.Param("service-id")
		if serviceId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing service-id"))
			return
		}
		serviceUUID, err := uuid.Parse(serviceId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("nurse-id invalid (not a uuid)"))
			return
		}

		dtos, err := s.query.GetBySerivceId.Handle(ctx, serviceUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, dtos)
	}
}
