package nursehttpservice

import (
	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursecommands "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		map service with nursing
//	@Description	map service with nursing
//	@Tags			nurse
//	@Accept			json
//	@Produce		json
//	@Param			nurse-id	path		string					true								"Nursing ID (UUID)"
//	@Param			create		form		body					nursecommands.MapNursingServiceDTO	true	"account creation data"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/nurses/{nurse-id}/services [post]
//	@Security		ApiKeyAuth
func (s *nurseHttpService) handleMapNursingServices() gin.HandlerFunc {
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

		var dto nursecommands.MapNursingServiceDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.MapNurseService.Handle(ctx.Request.Context(), nurseUUID, dto.ServiceIds); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
