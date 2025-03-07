package nursehttpservice

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
)

// @Summary		get nursing with filter
// @Description	get nursing with filter
// @Tags			nurse
// @Accept			json
// @Produce		json
// @Param			service-id	query		string					false	"services id"
// @Param			nurse-name	query		string					false	"nursing name"
// @Param			rate		query		string					false	"rate"
// @Param			page		query		string					false	"page"
// @Param			size		query		string					false	"size"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/nurses [get]
// @Security		ApiKeyAuth
func (s *nurseHttpService) handleGetNursingsWithFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errParseQuery error
		serviceId := ctx.Query("service-id")
		if serviceId != "" {
			_, errParseQuery = uuid.Parse(serviceId)
			if errParseQuery != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a uuid)"))
				return
			}
		}
		nurseName := ctx.Query("nurse-name")
		fmt.Println("nurseName: ", nurseName)

		rate := ctx.Query("rate")
		var rateFloat float64
		if rate != "" {
			rateFloat, errParseQuery = strconv.ParseFloat(strings.TrimSpace(rate), 64)
			if errParseQuery != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("rate must be a number (float)"))
				return
			}
		}

		paramFilter := &nursequeries.NurseFilterDTO{
			ServiceId: serviceId,
			NurseName: nurseName,
			Rate:      rateFloat,
		}

		page := ctx.Query("page")
		pageInt, _ := strconv.Atoi(page)
		size := ctx.Query("size")
		sizeInt, _ := strconv.Atoi(size)
		paramPaging := &common.Paging{
			Page:  pageInt,
			Size:  sizeInt,
			Total: 0,
		}

		param := nursequeries.NurseRequestQueryDTO{
			Filter: paramFilter,
			Paging: paramPaging,
		}

		result, err := s.query.GetWithFilter.Handle(ctx, &param)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseGetWithPagination(ctx, result, param.Paging, param.Filter)
	}
}
