package repl_auction

import (
	"context"
	"encoding/json"
	"event-trace/internal/busi/core/instancejob/common"

	"math/big"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	"github.com/umbracle/ethgo"
	ethabi "github.com/umbracle/ethgo/abi"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type ReplAuction struct {
	EventName string
}

const (
	FILReceivedEventHash = "0xa8460747e40836bec1c1362c948913239b816f81cc759fe9d7597176fa4ac648"
	FILReceivedEventName = "FILReceived(uint256,address)"

	TxnType         = 0
	InternalTXNType = 1
)

func NewInstance() ReplAuction {
	return ReplAuction{"ReplAuction"}
}

func (replAuction ReplAuction) GetEventName() string {
	return replAuction.EventName
}

func (replAuction ReplAuction) EventTracing(ctx context.Context, node *api.FullNodeStruct, args ...string) error {
	g, ctx := errgroup.WithContext(ctx)

	// EOA -> replAuction contract
	g.Go(func() error {
		return replAuction.tracingReplAuctionEventTXNCron(ctx, node, args[0], FILReceivedEventHash, FILReceivedEventName)
	})

	// CA -> replAuction contract
	g.Go(func() error {
		return replAuction.tracingReplAuctionEventCronInInternalTXN(ctx, node, args[0], FILReceivedEventHash, FILReceivedEventName)
	})

	return g.Wait()
}

// CA -> replAuction contract(EOA -> middle contracts -> replAuction contract, internal transaction)
func (replAuction ReplAuction) tracingReplAuctionEventCronInInternalTXN(ctx context.Context, _ *api.FullNodeStruct, replAuctionAddress, eventHash, eventName string) error {
	return common.TracingContractEventCronInInternalTXN(ctx, replAuctionAddress, eventHash, eventName, getTheEventContent)
}

// EOA -> replAuction contract(transaction)
func (replAuction ReplAuction) tracingReplAuctionEventTXNCron(ctx context.Context, _ *api.FullNodeStruct, replAuctionAddress, eventHash, eventName string) error {
	return common.TracingContractEventTXNCron(ctx, replAuctionAddress, eventHash, eventName, false, getTheEventContent)
}

type FILReceivedObj struct {
	A *big.Int
	B ethgo.Address
}

func getTheEventContent(eventName string, ethLog *ethtypes.EthLog) string {
	switch eventName {
	case FILReceivedEventName:
		typ := ethabi.MustNewType("tuple(uint256 a, address b)")
		var output FILReceivedObj
		err := typ.DecodeStruct(ethLog.Data, &output)
		if err != nil {
			log.Errorf("Get Error during decoding the FILReceived: %v", err)
			return ""
		}

		onNewFundReceived := FILReceived{
			Amount:    output.A.String(),
			AgentAddr: output.B.String(),
		}

		data, _ := json.Marshal(onNewFundReceived)
		return string(data)
	}
	return ""
}
