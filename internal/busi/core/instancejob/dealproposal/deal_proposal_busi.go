package dealproposal

import (
	"busi/pkg/models/fevm"
	"busi/pkg/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	mbig "math/big"

	log "github.com/sirupsen/logrus"
)

type DealProposalEventTracingCronFn func(context.Context, *api.FullNodeStruct) error

// 0xfd6419d07e4c269e58d0c63969756c2124155b4a8d6dd08b8cd46e3a9acbf625 is event - DealProposalCreate(bytes32,uint64,bool,uint256)'s hash
const (
	DealProposalCreateEventHash = "0xfd6419d07e4c269e58d0c63969756c2124155b4a8d6dd08b8cd46e3a9acbf625"
	DealProposalCreateEventName = "DealProposalCreate(bytes32,uint64,bool,uint256)"
)

func TracingDealProposalEventCron(ctx context.Context, node *api.FullNodeStruct) error {
	var (
		maxHeightEvmReceipt fevm.EVMReceipt
		recordedHeight      fevm.EventHeightCheckpoint
	)
	evmReceipts := make([]*fevm.EVMReceipt, 0)

	recordedHeight.EventHash = DealProposalCreateEventHash
	recordedHeight.EventName = DealProposalCreateEventName

	// select * from event_height_checkpoint where event_hash = DealProposalCreateHash;
	// select max(height) from evm_receipt;
	// select * from evm_receipt where height > leftHeight and height < rightHeight and logs like '%DealProposalCreateHash%' order by height desc;
	if _, err := utils.X.Where("event_hash = ?", DealProposalCreateEventHash).Get(&recordedHeight); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if _, err := utils.X.Select("max(height) as height").Get(&maxHeightEvmReceipt); err != nil {
		log.Errorf("execute sql error: %v", err)
		return err
	}

	if err := utils.X.Where("height between ? and ? and logs like ?", recordedHeight.MaxRecordedHeight+1, maxHeightEvmReceipt.Height+1, "%"+DealProposalCreateEventHash+"%").Asc("height").Find(&evmReceipts); err != nil {
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
			if ethLog.Topics[0].String() != DealProposalCreateEventHash {
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
				EventName:       DealProposalCreateEventName,
			}

			// invoke getDealProposal on chain
			fromEthAddr, err := ethtypes.ParseEthAddress(receipt.From) // any eth address is ok.
			if err != nil {
				log.Errorf("parsing `from` eth address failed: %v", err)
				continue
			}
			res, err := getDealProposal(ctx, node, receipt.To, ethLog.Topics[1].String(), fromEthAddr)
			if err != nil {
				log.Errorf("eth call for deal proposal failed: %v", err)
				continue
			}

			var dpc market.DealProposal
			if err := dpc.UnmarshalCBOR(bytes.NewReader(res)); err != nil {
				log.Errorf("cbor unmarshal failed: %v", err)
				continue
			}

			res, _ = json.Marshal(&dpc)
			fevmEvent.Note = string(res)

			if _, err := utils.X.Insert(&fevmEvent); err != nil {
				log.Errorf("execute sql error: %v", err)
				return err
			}

			// update or insert event_height_checkoutpoint
			recordedHeight.MaxRecordedHeight = uint64(receipt.Height)
			if err := updateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
				return err
			}
		}
	}

	recordedHeight.MaxRecordedHeight = uint64(maxHeightEvmReceipt.Height)
	if err := updateEventHeightCheckoutpoint(ctx, &recordedHeight); err != nil {
		return err
	}

	return nil
}

func TracingDealProposalEvent(ctx context.Context, node *api.FullNodeStruct, minHeight, maxHeight uint64) error {
	evmReceipts := make([]*fevm.EVMReceipt, 0)

	if err := utils.X.Where("height between ? and ? and logs like ?", minHeight, maxHeight, "%"+DealProposalCreateEventHash+"%").Asc("height").Find(&evmReceipts); err != nil {
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
			if ethLog.Topics[0].String() != DealProposalCreateEventHash {
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
				EventName:       DealProposalCreateEventName,
			}

			// invoke getDealProposal on chain
			fromEthAddr, err := ethtypes.ParseEthAddress(receipt.From) // any eth address is ok.
			if err != nil {
				log.Errorf("parsing `from` eth address failed: %v", err)
				continue
			}
			res, err := getDealProposal(ctx, node, receipt.To, ethLog.Topics[1].String(), fromEthAddr)
			if err != nil {
				log.Errorf("eth call for deal proposal failed: %v", err)
				continue
			}

			var dpc market.DealProposal
			if err := dpc.UnmarshalCBOR(bytes.NewReader(res)); err != nil {
				log.Errorf("cbor unmarshal failed: %v", err)
				continue
			}

			res, _ = json.Marshal(&dpc)
			fevmEvent.Note = string(res)

			if _, err := utils.X.Insert(&fevmEvent); err != nil {
				log.Errorf("execute sql error: %v", err)
				return err
			}
		}
	}

	return nil
}

func updateEventHeightCheckoutpoint(ctx context.Context, ec *fevm.EventHeightCheckpoint) error {
	b, err := utils.X.ID(ec.Id).Update(ec)
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

// refer - https://github.com/filecoin-project/boost/blob/main/storagemarket/contract_deal_monitor.go#L27
func getDealProposal(ctx context.Context, node *api.FullNodeStruct, topicContractAddress string, topicDealProposalID string, fromEthAddr ethtypes.EthAddress) ([]byte, error) {
	// GetDealProposal is a free data retrieval call binding the contract method 0xf4b2e4d8.
	_params := "0xf4b2e4d8" + topicDealProposalID[2:] // cut 0x prefix

	toEthAddr, err := ethtypes.ParseEthAddress(topicContractAddress)
	if err != nil {
		return nil, fmt.Errorf("parsing `to` eth address failed: %w", err)
	}

	params, err := ethtypes.DecodeHexString(_params)
	if err != nil {
		return nil, fmt.Errorf("decoding params failed: %w", err)
	}

	latest := "latest"
	blkParam := ethtypes.EthBlockNumberOrHash{
		PredefinedBlock: &latest,
	}

	res, err := node.EthCall(ctx, ethtypes.EthCall{
		From: &fromEthAddr,
		To:   &toEthAddr,
		Data: params,
	}, blkParam)
	if err != nil {
		return nil, fmt.Errorf("eth call erred: %w", err)
	}

	begin, length, err := lengthPrefixPointsTo(res)
	if err != nil {
		return nil, fmt.Errorf("length prefix points erred: %w", err)
	}

	return res[begin : begin+length], nil
}

func lengthPrefixPointsTo(output []byte) (int, int, error) {
	index := 0
	boffset := mbig.NewInt(0).SetBytes(output[index : index+32])
	boffset.Add(boffset, mbig.NewInt(32))
	boutputLen := mbig.NewInt(int64(len(output)))

	if boffset.Cmp(boutputLen) > 0 {
		return 0, 0, fmt.Errorf("offset %v is over boundary; len: %v", boffset, boutputLen)
	}

	if boffset.BitLen() > 63 {
		return 0, 0, fmt.Errorf("offset larger than int64: %v", boffset)
	}

	offset := int(boffset.Uint64())
	lengthBig := mbig.NewInt(0).SetBytes(output[offset-32 : offset])

	size := mbig.NewInt(0)
	size.Add(size, boffset)
	size.Add(size, lengthBig)
	if size.BitLen() > 63 {
		return 0, 0, fmt.Errorf("len larger than int64: %v", size)
	}

	if size.Cmp(boutputLen) > 0 {
		return 0, 0, fmt.Errorf("length insufficient %v require %v", boutputLen, size)
	}

	return int(boffset.Uint64()), int(lengthBig.Uint64()), nil
}

// case 2:
// select * from evm_receipt where height bewteen xxx and xxx and logs like '%DealProposalCreateHash%' order by height desc;
