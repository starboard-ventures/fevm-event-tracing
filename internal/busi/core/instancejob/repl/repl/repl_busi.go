package repl

import (
	"context"
	"encoding/json"
	"event-trace/internal/busi/core/instancejob/common"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/umbracle/ethgo"

	"math/big"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	ethabi "github.com/umbracle/ethgo/abi"
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
	return common.TracingContractEventTXNCron(ctx, replAddress, eventHash, eventName, false, getTheEventContent)
}

type OnNewFundReceivedObj struct {
	A *big.Int
	B ethgo.Address
}

func getTheEventContent(eventName string, ethLog *ethtypes.EthLog) string {
	switch eventName {
	case ReplTransferEventName:
		typ := ethabi.MustNewType("tuple(uint256 a, address b)")
		var output OnNewFundReceivedObj
		err := typ.DecodeStruct(ethLog.Data, &output)
		if err != nil {
			log.Errorf("Get Error during decoding the FILReceived: %v", err)
			return ""
		}
		onNewFundReceived := OnNewFundReceived{
			Amount:    output.A.String(),
			AgentAddr: output.B.String(),
		}

		data, _ := json.Marshal(onNewFundReceived)
		return string(data)
	}
	return ""
}
