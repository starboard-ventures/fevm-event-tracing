package wfil

import (
	"context"
	"encoding/json"
	"event-trace/internal/busi/core/instancejob/common"
	"event-trace/pkg/models/fevm"
	"event-trace/pkg/utils"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	log "github.com/sirupsen/logrus"
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
)

func NewInstance() Wfil {
	return Wfil{"Wfil"}
}

func (dpc Wfil) GetEventName() string {
	return dpc.EventName
}

func (dpc Wfil) EventTracing(ctx context.Context, node *api.FullNodeStruct, args ...string) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return dpc.tracingWfilEventCron(ctx, node, args[0], WfilDepositEventHash, WfilDepositEventName)
	})

	g.Go(func() error {
		return dpc.tracingWfilEventCron(ctx, node, args[0], WfilWithdrawalEventHash, WfilWithdrawalEventName)
	})

	return g.Wait()
}

func (dpc Wfil) tracingWfilEventCron(ctx context.Context, node *api.FullNodeStruct, wfilAddress, eventHash, eventName string) error {
	var (
		maxHeightEvmReceipt fevm.EVMReceipt
		recordedHeight      fevm.EventHeightCheckpoint
	)
	evmReceipts := make([]*fevm.EVMReceipt, 0)

	recordedHeight.EventHash = eventHash
	recordedHeight.EventName = eventName

	// select * from event_height_checkpoint where event_hash = %eventHash%;
	// select max(height) from evm_receipt;
	// select * from evm_receipt where height > leftHeight and height < rightHeight and logs like '%eventHash%' order by height desc;
	if _, err := utils.X.Where("event_hash = ?", eventHash).Get(&recordedHeight); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if _, err := utils.X.Select("max(height) as height").Get(&maxHeightEvmReceipt); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if err := utils.X.Where("height between ? and ? and \"to\" = ? and logs like ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height+1, wfilAddress, "%"+eventHash+"%").Asc("height").Find(&evmReceipts); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	for _, receipt := range evmReceipts {
		logs := make([]ethtypes.EthLog, 0)
		if err := json.Unmarshal([]byte(receipt.Logs), &logs); err != nil {
			log.Warnf("Unmarshal receipt[height: %v] log err: %v", receipt.Height, err)
			continue
		}

		for _, ethLog := range logs {
			if ethLog.Topics[0].String() != eventHash {
				continue
			}

			fevmEvent := fevm.FevmEvent{
				Height:          uint64(receipt.Height),
				TransactionHash: receipt.TransactionHash,
				From:            receipt.From,
				To:              receipt.To,
				Status:          receipt.Status,
				LogsBloom:       receipt.LogsBloom,
				Logs:            receipt.Logs,
				EventHash:       ethLog.Topics[0].String(),
				EventName:       eventName,
			}

			if _, err := utils.X.Insert(&fevmEvent); err != nil {
				log.Errorf("execute sql error: %v", err)
				return err
			}

			// update or insert event_height_checkoutpoint
			recordedHeight.MaxRecordedHeight = uint64(receipt.Height)
			if err := common.UpdateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
				return err
			}
		}
	}

	recordedHeight.MaxRecordedHeight = uint64(maxHeightEvmReceipt.Height)
	if err := common.UpdateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
		return err
	}

	return nil
}
