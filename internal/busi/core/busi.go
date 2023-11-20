package core

import (
	"context"
	"event-trace/internal/busi/core/instancejob"
	"event-trace/internal/busi/core/instancejob/dealproposal"
	"event-trace/internal/busi/core/instancejob/wfil"
	"event-trace/pkg/utils"
	"net/http"
)

func DealProposalCreateEventCronHandle(ctx context.Context, lotus0 string) *utils.BuErrorResponse {
	lotusAPI, closer, err := utils.LotusHandshake(ctx, lotus0)
	if err != nil {
		return &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
	}
	defer closer()

	// Cronjob by dolphin scheduler
	if err := instancejob.NewCronJob(lotusAPI, 0, 0, dealproposal.NewInstance()).TracingJobExecute(ctx); err != nil {
		return &utils.BuErrorResponse{HttpCode: http.StatusOK, Response: &utils.Response{Code: utils.CodeInternalServer, Message: err.Error()}}
	}

	return nil
}

func DealProposalCreateEventHandle(ctx context.Context, r *RequestHeight) *utils.BuErrorResponse {
	lotusAPI, closer, err := utils.LotusHandshake(ctx, r.Lotus0)
	if err != nil {
		return &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
	}
	defer closer()

	// manual call
	if err := dealproposal.NewInstance().TracingDealProposalEvent(ctx, lotusAPI, r.MinHeight, r.MaxHeight); err != nil {
		return &utils.BuErrorResponse{HttpCode: http.StatusOK, Response: &utils.Response{Code: utils.CodeInternalServer, Message: err.Error()}}
	}

	return nil
}

func WfilEventCronHandle(ctx context.Context, lotus0, wfilContract string) *utils.BuErrorResponse {
	lotusAPI, closer, err := utils.LotusHandshake(ctx, lotus0)
	if err != nil {
		return &utils.BuErrorResponse{HttpCode: http.StatusInternalServerError, Response: utils.ErrInternalServer}
	}
	defer closer()

	// Cronjob by dolphin scheduler
	if err := instancejob.NewCronJob(lotusAPI, 0, 0, wfil.NewInstance()).TracingJobExecute(ctx, wfilContract); err != nil {
		return &utils.BuErrorResponse{HttpCode: http.StatusOK, Response: &utils.Response{Code: utils.CodeInternalServer, Message: err.Error()}}
	}

	return nil
}
