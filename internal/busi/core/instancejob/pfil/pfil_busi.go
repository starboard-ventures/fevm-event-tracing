package pfil

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
	var (
		maxHeightEvmReceipt fevm.EVMReceipt
		recordedHeight      fevm.EventHeightCheckpoint
	)
	evmInternalTxn := make([]*fevm.EVMInternalTXN, 0)

	recordedHeight.EventHash = eventHash
	recordedHeight.EventName = eventName
	recordedHeight.TxnType = InternalTXNType

	if _, err := utils.X.Where("event_hash = ? and txn_type = ?", eventHash, InternalTXNType).Get(&recordedHeight); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if _, err := utils.X.Select("max(height) as height").Get(&maxHeightEvmReceipt); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if err := utils.X.Where("height between ? and ? and \"to\" = ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height+1, pfilAddress).Asc("height").Find(&evmInternalTxn); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	for _, internalTXN := range evmInternalTxn {
		var (
			b          bool
			err        error
			receiptTmp fevm.EVMReceipt
		)

		b, err = utils.X.Where("height = ? and transaction_hash = ?", internalTXN.Height, internalTXN.ParentHash).Get(&receiptTmp)
		if err != nil {
			log.Errorf("execute sql error: %v", err)
			return err
		}

		if !b {
			log.Warnf("can't find internal transaction's transaction, height: %v, hash: %v", internalTXN.Height, internalTXN.ParentHash)
			continue
		}

		logs := make([]ethtypes.EthLog, 0)
		if err := json.Unmarshal([]byte(receiptTmp.Logs), &logs); err != nil {
			log.Warnf("Unmarshal receipt[height: %v] log err: %v", receiptTmp.Height, err)
			continue
		}

		for _, ethLog := range logs {
			if ethLog.Address.String() != pfilAddress || ethLog.Topics[0].String() != eventHash {
				continue
			}

			fevmEvent := fevm.FevmEvent{
				ContractAddress: pfilAddress,
				Height:          uint64(receiptTmp.Height),
				TransactionHash: receiptTmp.TransactionHash,
				From:            receiptTmp.From,
				To:              receiptTmp.To,
				Status:          receiptTmp.Status,
				LogsBloom:       receiptTmp.LogsBloom,
				Logs:            receiptTmp.Logs,
				EventHash:       ethLog.Topics[0].String(),
				EventName:       eventName,
			}

			fevmEvent.Note = pfil.getTheEventContent(eventName, ethLog.Topics[1].String(), ethLog.Topics[2].String(), ethLog.Data.String())

			if _, err := utils.X.Insert(&fevmEvent); err != nil {
				log.Errorf("execute sql error: %v", err)
				return err
			}

			// update or insert event_height_checkoutpoint
			recordedHeight.MaxRecordedHeight = uint64(receiptTmp.Height)
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

// EOA -> pfil contract(transaction)
func (pfil Pfil) tracingPfilEventTXNCron(ctx context.Context, _ *api.FullNodeStruct, pfilAddress, eventHash, eventName string) error {
	var (
		maxHeightEvmReceipt fevm.EVMReceipt
		recordedHeight      fevm.EventHeightCheckpoint
	)
	evmReceipts := make([]*fevm.EVMReceipt, 0)

	recordedHeight.EventHash = eventHash
	recordedHeight.EventName = eventName
	recordedHeight.TxnType = TxnType

	// select * from event_height_checkpoint where event_hash = %eventHash%;
	// select max(height) from evm_receipt;
	// select * from evm_receipt where height > leftHeight and height < rightHeight and logs like '%eventHash%' order by height desc;
	if _, err := utils.X.Where("event_hash = ? and txn_type = ?", eventHash, TxnType).Get(&recordedHeight); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if _, err := utils.X.Select("max(height) as height").Get(&maxHeightEvmReceipt); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if err := utils.X.Where("height between ? and ? and \"to\" = ? and logs like ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height+1, pfilAddress, "%"+eventHash+"%").Asc("height").Find(&evmReceipts); err != nil {
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
				ContractAddress: pfilAddress,
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

			fevmEvent.Note = pfil.getTheEventContent(eventName, ethLog.Topics[1].String(), ethLog.Topics[2].String(), ethLog.Data.String())

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

func (pfil Pfil) getTheEventContent(eventName string, eventIndex1, eventIndex2, eventData string) string {
	switch eventName {
	case PfilTransferEventName:
		transfer := Transfer{
			From:   eventIndex1,
			To:     eventIndex2,
			Amount: eventData,
		}

		data, _ := json.Marshal(transfer)
		return string(data)
	}
	return ""
}
