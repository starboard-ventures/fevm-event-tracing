package wfil

import (
	"context"
	"encoding/json"
	"event-trace/internal/busi/core/instancejob/common"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	"golang.org/x/sync/errgroup"
)

type Wfil struct {
	EventName string
}

const (
	WfilDepositEventHash = "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"
	WfilDepositEventName = "Deposit(address,uint256)"

	WfilWithdrawalEventHash = "0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"
	WfilWithdrawalEventName = "Withdrawal(address,uint256)"

	TxnType         = 0
	InternalTXNType = 1
)

func NewInstance() Wfil {
	return Wfil{"Wfil"}
}

func (wfil Wfil) GetEventName() string {
	return wfil.EventName
}

func (wfil Wfil) EventTracing(ctx context.Context, node *api.FullNodeStruct, args ...string) error {
	g, ctx := errgroup.WithContext(ctx)

	// EOA -> wfil contract
	g.Go(func() error {
		return wfil.tracingWfilEventTXNCron(ctx, node, args[0], WfilDepositEventHash, WfilDepositEventName)
	})

	g.Go(func() error {
		return wfil.tracingWfilEventTXNCron(ctx, node, args[0], WfilWithdrawalEventHash, WfilWithdrawalEventName)
	})

	// CA -> wfil contract
	g.Go(func() error {
		return wfil.tracingWfilEventCronInInternalTXN(ctx, node, args[0], WfilDepositEventHash, WfilDepositEventName)
	})

	g.Go(func() error {
		return wfil.tracingWfilEventCronInInternalTXN(ctx, node, args[0], WfilWithdrawalEventHash, WfilWithdrawalEventName)
	})

	return g.Wait()
}

// CA -> wfil contract(EOA -> middle contracts -> wfil contract, internal transaction)
func (wfil Wfil) tracingWfilEventCronInInternalTXN(ctx context.Context, _ *api.FullNodeStruct, wfilAddress, eventHash, eventName string) error {
	return common.TracingContractEventCronInInternalTXN(ctx, wfilAddress, eventHash, eventName, getTheEventContent)
}

// EOA -> wfil contract(transaction)
func (wfil Wfil) tracingWfilEventTXNCron(ctx context.Context, _ *api.FullNodeStruct, wfilAddress, eventHash, eventName string) error {
	return common.TracingContractEventTXNCron(ctx, wfilAddress, eventHash, eventName, true, getTheEventContent)
}

func getTheEventContent(eventName string, ethLog *ethtypes.EthLog) string {
	switch eventName {
	case WfilDepositEventName:
		deposit := Deposit{
			From:   ethLog.Topics[1].String(),
			Amount: ethLog.Data.String(),
		}

		data, _ := json.Marshal(deposit)
		return string(data)
	case WfilWithdrawalEventName:
		withdrawal := Withdrawal{
			To:     ethLog.Topics[1].String(),
			Amount: ethLog.Data.String(),
		}

		data, _ := json.Marshal(withdrawal)
		return string(data)
	}
	return ""
}
