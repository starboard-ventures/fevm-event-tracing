package repl

import (
	"context"
	"encoding/json"
	"event-trace/internal/busi/core/instancejob/common"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	"golang.org/x/sync/errgroup"
)

type Repl struct {
	EventName string
}

const (
	ReplTransferEventHash = "0x8e98551cebbb96f1784e350829af95b8b3bddc33df724ef2cde19d511ac68337"
	ReplTransferEventName = "OnNewFundReceived(uint256,address)"

	TxnType         = 0
	InternalTXNType = 1
)

func NewInstance() Repl {
	return Repl{"Repl"}
}

func (repl Repl) GetEventName() string {
	return repl.EventName
}

func (repl Repl) EventTracing(ctx context.Context, node *api.FullNodeStruct, args ...string) error {
	g, ctx := errgroup.WithContext(ctx)

	// EOA -> repl contract
	g.Go(func() error {
		return repl.tracingReplEventTXNCron(ctx, node, args[0], ReplTransferEventHash, ReplTransferEventName)
	})

	// CA -> repl contract
	g.Go(func() error {
		return repl.tracingReplEventCronInInternalTXN(ctx, node, args[0], ReplTransferEventHash, ReplTransferEventName)
	})

	return g.Wait()
}

// CA -> repl contract(EOA -> middle contracts -> repl contract, internal transaction)
func (repl Repl) tracingReplEventCronInInternalTXN(ctx context.Context, _ *api.FullNodeStruct, replAddress, eventHash, eventName string) error {
	return common.TracingContractEventCronInInternalTXN(ctx, replAddress, eventHash, eventName, getTheEventContent)
}

// EOA -> repl contract(transaction)
func (repl Repl) tracingReplEventTXNCron(ctx context.Context, _ *api.FullNodeStruct, replAddress, eventHash, eventName string) error {
	return common.TracingContractEventTXNCron(ctx, replAddress, eventHash, eventName, getTheEventContent)
}

func getTheEventContent(eventName string, ethLog *ethtypes.EthLog) string {
	switch eventName {
	case ReplTransferEventName:
		onNewFundReceived := OnNewFundReceived{
			// Amount:    ,
			// AgentAddr: ,
		}

		data, _ := json.Marshal(onNewFundReceived)
		return string(data)
	}
	return ""
}
