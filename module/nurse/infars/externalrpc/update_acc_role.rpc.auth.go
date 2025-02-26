package nurseexternalrpc

// func (ex *externalAccountService) UpdateRoleNurseToStaffRPC(ctx context.Context, accId uuid.UUID, payload *nursecommands.UpdateRoleRequestRPC) error {
// 	token, ok := ctx.Value(common.KeyToken).(string)
// 	if !ok {
// 		return common.NewInternalServerError().
// 			WithReason("cannot get delete account profile").WithInner("missing token to fetching data from other service")
// 	}
//
// 	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
// 		Method:  "PATCH",
// 		URL:     ex.apiURL + "/external/rpc/accounts/" + accId.String() + "/role",
// 		Token:   token,
// 		Payload: payload,
// 	})
// 	if err != nil {
// 		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
// 		return resp
// 	}
//
// 	success, ok := response["success"].(bool)
// 	if !ok || !success {
// 		return common.ExtractErrorFromResponse(response)
// 	}
//
// 	return nil
// }
