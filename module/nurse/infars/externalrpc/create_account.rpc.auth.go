package nurseexternalrpc

import (
	"context"

	"github.com/PhuPhuoc/curanest-nurse-service/common"
	nursecommands "github.com/PhuPhuoc/curanest-nurse-service/module/nurse/usecase/commands"
)

func (ex *externalAccountService) CreateAccountRPC(
	ctx context.Context,
	entity *nursecommands.AccountInfoDTO,
) (*nursecommands.ResponseCreateAccountDTO, error) {
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     ex.apiURL + "/external/rpc/accounts",
		Payload: entity,
	})
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot create accounts").WithInner("cannot call external api - " + err.Error())
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractDataFromResponse[nursecommands.ResponseCreateAccountDTO](response, "data")
}
