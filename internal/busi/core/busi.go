package core

import (
	"busi/pkg/utils"
	"context"
)

func EventHandle(ctx context.Context, r *RequestHeight) *utils.BuErrorResponse {
	// var (
	// 	err       error
	// 	cachemvkv *MailVerifyKV
	// )

	// mvkv := &MailVerifyKV{Email: r.Email, Code: r.Code, BrowserFingerPrintMetadata: r.BrowserFingerPrintMetadata}
	// if cachemvkv, err = mvkv.MailKVGet(ctx); err != nil {
	// 	return &utils.BuErrorResponse{http.StatusInternalServerError, utils.ErrAccountInternal}
	// }
	// if cachemvkv == nil {
	// 	return &utils.BuErrorResponse{http.StatusOK, utils.ErrMailCodeNotSend}
	// }

	// if mvkv.BrowserFingerPrintMetadata != cachemvkv.BrowserFingerPrintMetadata {
	// 	return &utils.BuErrorResponse{http.StatusOK, utils.ErrAccountCodeOrFingerPrint}
	// }
	// if mvkv.Code != cachemvkv.Code {
	// 	return &utils.BuErrorResponse{http.StatusOK, utils.ErrMailCodeVerifiedFailed}
	// }

	// if _, err = mvkv.MailKVSetVerifiedState(ctx, true); err != nil {
	// 	return &utils.BuErrorResponse{http.StatusInternalServerError, utils.ErrAccountInternal}
	// }
	return nil
}
