package nurseexternalrpc

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursequeries "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/queries"
)

func (ex *externalAccountService) GetAccountByIdRPC(ctx context.Context, accId string) (*nursequeries.ResponseAccountDTO, error) {
	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return nil, common.NewInternalServerError().
			WithReason("cannot get accounts profile").WithInner("missing token to fetching data from other service")
	}

	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "GET",
		URL:    ex.apiURL + "/external/rpc/accounts/" + accId,
		Token:  token,
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return nil, resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractDataFromResponse[nursequeries.ResponseAccountDTO](response, "data")
}
