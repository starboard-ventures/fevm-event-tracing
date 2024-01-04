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
	var (
		maxHeightEvmReceipt fevm.EVMReceipt
		recordedHeight      fevm.EventHeightCheckpoint
	)
	evmInternalTxn := make([]*fevm.EVMInternalTXN, 0)

	recordedHeight.EventHash = eventHash
	recordedHeight.EventName = eventName
	recordedHeight.TxnType = InternalTXNType

	if _, err := utils.X.Where("event_hash = ? and type = ?", eventHash, InternalTXNType).Get(&recordedHeight); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if _, err := utils.X.Select("max(height) as height").Get(&maxHeightEvmReceipt); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if err := utils.X.Where("height between ? and ? and \"to\" = ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height+1, wfilAddress).Asc("height").Find(&evmInternalTxn); err != nil {
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
			if ethLog.Address.String() != wfilAddress || ethLog.Topics[0].String() != eventHash {
				continue
			}

			fevmEvent := fevm.FevmEvent{
				ContractAddress: wfilAddress,
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

			fevmEvent.Note = wfil.getTheEventContent(eventName, ethLog.Topics[1].String(), ethLog.Data.String())

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

// EOA -> wfil contract(transaction)
func (wfil Wfil) tracingWfilEventTXNCron(ctx context.Context, _ *api.FullNodeStruct, wfilAddress, eventHash, eventName string) error {
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
	if _, err := utils.X.Where("event_hash = ? and type = ?", eventHash, TxnType).Get(&recordedHeight); err != nil {
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
				ContractAddress: wfilAddress,
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

			fevmEvent.Note = wfil.getTheEventContent(eventName, ethLog.Topics[1].String(), ethLog.Data.String())

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

func (wfil Wfil) getTheEventContent(eventName string, eventIndex, eventData string) string {
	switch eventName {
	case WfilDepositEventName:
		deposit := Deposit{
			From:   eventIndex,
			Amount: eventData,
		}

		data, _ := json.Marshal(deposit)
		return string(data)
	case WfilWithdrawalEventName:
		withdrawal := Withdrawal{
			To:     eventIndex,
			Amount: eventData,
		}

		data, _ := json.Marshal(withdrawal)
		return string(data)
	}
	return ""
}

func (wfil Wfil) tracingWfilEvent(ctx context.Context, _ *api.FullNodeStruct, wfilAddress, eventHash, eventName string, minHeight, maxHeight uint64) error {
	evmReceipts := make([]*fevm.EVMReceipt, 0)

	if err := utils.X.Where("height between ? and ? and \"to\" = ? and logs like ?", minHeight, maxHeight, wfilAddress, "%"+eventHash+"%").Asc("height").Find(&evmReceipts); err != nil {
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

			fevmEvent.Note = wfil.getTheEventContent(eventName, ethLog.Topics[1].String(), ethLog.Data.String())

			if _, err := utils.X.Insert(&fevmEvent); err != nil {
				log.Errorf("execute sql error: %v", err)
				return err
			}
		}
	}

	return nil
}

func (wfil Wfil) TracingWfilEvent(ctx context.Context, _ *api.FullNodeStruct, minHeight, maxHeight uint64, wfilAddress string) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return wfil.tracingWfilEvent(ctx, nil, wfilAddress, WfilDepositEventHash, WfilDepositEventName, minHeight, maxHeight)
	})

	g.Go(func() error {
		return wfil.tracingWfilEvent(ctx, nil, wfilAddress, WfilWithdrawalEventHash, WfilWithdrawalEventName, minHeight, maxHeight)
	})

	return g.Wait()
}
