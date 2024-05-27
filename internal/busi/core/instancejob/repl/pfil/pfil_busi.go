package pfil

import (
	"context"
	"encoding/json"
	"event-trace/internal/busi/core/instancejob/common"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	"golang.org/x/sync/errgroup"
)

type Pfil struct {
	EventName string
}

const (
	PfilTransferEventHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	PfilTransferEventName = "Transfer(address,address,uint256)"

	TxnType         = 0
	InternalTXNType = 1
)

func NewInstance() Pfil {
	return Pfil{"Pfil"}
}

func (pfil Pfil) GetEventName() string {
	return pfil.EventName
}

func (pfil Pfil) EventTracing(ctx context.Context, node *api.FullNodeStruct, args ...string) error {
	g, ctx := errgroup.WithContext(ctx)

	// EOA -> pfil contract
	g.Go(func() error {
		return pfil.tracingPfilEventTXNCron(ctx, node, args[0], PfilTransferEventHash, PfilTransferEventName)
	})

	// CA -> pfil contract
	g.Go(func() error {
		return pfil.tracingPfilEventCronInInternalTXN(ctx, node, args[0], PfilTransferEventHash, PfilTransferEventName)
	})

	return g.Wait()
}

// CA -> pfil contract(EOA -> middle contracts -> pfil contract, internal transaction)
func (pfil Pfil) tracingPfilEventCronInInternalTXN(ctx context.Context, _ *api.FullNodeStruct, pfilAddress, eventHash, eventName string) error {
	return common.TracingContractEventCronInInternalTXN(ctx, pfilAddress, eventHash, eventName, getTheEventContent)
}

// EOA -> pfil contract(transaction)
func (pfil Pfil) tracingPfilEventTXNCron(ctx context.Context, _ *api.FullNodeStruct, pfilAddress, eventHash, eventName string) error {
	return common.TracingContractEventTXNCron(ctx, pfilAddress, eventHash, eventName, true, getTheEventContent)
}

func getTheEventContent(eventName string, ethLog *ethtypes.EthLog) string {
	switch eventName {
	case PfilTransferEventName:
		transfer := Transfer{
			From:   ethLog.Topics[1].String(),
			To:     ethLog.Topics[2].String(),
			Amount: ethLog.Data.String(),
		}

		data, _ := json.Marshal(transfer)
		return string(data)
	}
	return ""
}
