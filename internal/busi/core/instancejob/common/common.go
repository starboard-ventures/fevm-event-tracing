package common

import (
	"context"
	"encoding/json"
	"event-trace/pkg/models/fevm"
	"event-trace/pkg/utils"

	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	log "github.com/sirupsen/logrus"
)

func UpdateEventHeightCheckoutpoint(ctx context.Context, ec *fevm.EventHeightCheckpoint) error {
	b, err := utils.X.ID(ec.Id).MustCols("txn_type").Update(ec)
	if err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if b == 0 { // there aren't any records in the table
		_, err = utils.X.InsertOne(ec)
		if err != nil {
			log.Errorf("execute sql error: %v", err)
			return err
		}
	}

	return nil
}

const (
	TxnType         = 0
	InternalTXNType = 1

	Finality = 900
)

type GetEventContentCallback func(string, *ethtypes.EthLog) string

// EOA -> x contract(transaction)
func TracingContractEventTXNCron(ctx context.Context, contractAddress, eventHash, eventName string, assignTo bool, callEventContentFn GetEventContentCallback) error {
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

	if maxHeightEvmReceipt.Height >= Finality {
		maxHeightEvmReceipt.Height -= Finality
	}

	if assignTo {
		if err := utils.X.Where("height between ? and ? and \"to\" = ? and logs like ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height, contractAddress, "%"+eventHash+"%").Asc("height").Find(&evmReceipts); err != nil {
			log.Errorf("execute sql error: %v", err)
			return err
		}
	} else {
		if err := utils.X.Where("height between ? and ? and logs like ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height, "%"+eventHash+"%").Asc("height").Find(&evmReceipts); err != nil {
			log.Errorf("execute sql error: %v", err)
			return err
		}
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
				ContractAddress: contractAddress,
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

			fevmEvent.Note = callEventContentFn(eventName, &ethLog)

			if _, err := utils.X.Insert(&fevmEvent); err != nil {
				log.Errorf("execute sql error: %v", err)
				return err
			}

			// update or insert event_height_checkoutpoint
			recordedHeight.MaxRecordedHeight = uint64(receipt.Height)
			if err := UpdateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
				return err
			}
		}
	}

	recordedHeight.MaxRecordedHeight = uint64(maxHeightEvmReceipt.Height)
	if err := UpdateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
		return err
	}

	return nil
}

// CA -> x contract(EOA -> middle contracts -> x contract, internal transaction)
func TracingContractEventCronInInternalTXN(ctx context.Context, contractAddress, eventHash, eventName string, callEventContentFn GetEventContentCallback) error {
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

	if maxHeightEvmReceipt.Height >= Finality {
		maxHeightEvmReceipt.Height -= Finality
	}

	if err := utils.X.Where("height between ? and ? and \"to\" = ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height, contractAddress).Asc("height").Find(&evmInternalTxn); err != nil {
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
			if ethLog.Address.String() != contractAddress || ethLog.Topics[0].String() != eventHash {
				continue
			}

			fevmEvent := fevm.FevmEvent{
				ContractAddress: contractAddress,
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

			fevmEvent.Note = callEventContentFn(eventName, &ethLog)

			if _, err := utils.X.Insert(&fevmEvent); err != nil {
				log.Errorf("execute sql error: %v", err)
				return err
			}

			// update or insert event_height_checkoutpoint
			recordedHeight.MaxRecordedHeight = uint64(receiptTmp.Height)
			if err := UpdateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
				return err
			}
		}
	}

	recordedHeight.MaxRecordedHeight = uint64(maxHeightEvmReceipt.Height)
	if err := UpdateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
		return err
	}

	return nil
}
