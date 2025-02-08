package nurseexternalrpc

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
)

func (ex *externalAccountService) HardDeleteAccountProfileRPC(ctx context.Context, accId string) error {
	fmt.Println("in HardDeleteAccountProfileRPC")
	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return common.NewInternalServerError().
			WithReason("cannot get delete account profile").WithInner("missing token to fetching data from other service")
	}

	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "DELETE",
		URL:    ex.apiURL + "/external/rpc/accounts/" + accId,
		Token:  token,
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return common.ExtractErrorFromResponse(response)
	}

	return nil
}
