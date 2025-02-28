package nurserpcservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		get staff info by ids
// @Description	get staff info by ids
// @Tags			rpc: nurse
// @Accept			json
// @Produce		json
// @Param			create	form		body					nursequeries.StaffIdsQueryDTO	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/external/rpc/nurses/by-ids [post]
// @Security		ApiKeyAuth
func (s *nursingRPCService) handleGetStaffInfoByIds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto nursequeries.StaffIdsQueryDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		dtos, err := s.query.GetStaffByIds.Handle(ctx, dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, dtos)
	}
}
