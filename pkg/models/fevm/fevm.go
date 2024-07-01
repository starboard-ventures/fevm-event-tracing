package fevm

import "time"

type FevmEvent struct {
	Id              uint64 `json:"-" xorm:"bigserial autoincr"`
	ContractAddress string `xorm:"varchar(255) index notnull default ''"`
	Height          uint64 `xorm:"bigint notnull"`
	TransactionHash string `xorm:"varchar(255) notnull pk"`
	From            string `xorm:"varchar(255) index notnull default ''"`
	To              string `xorm:"varchar(255) index notnull default ''"`
	Status          int64  `xorm:"integer notnull default 0"`
	LogsBloom       string `xorm:"text notnull default ''"`
	Logs            string `xorm:"text notnull default ''"`
	EventHash       string `xorm:"varchar(255) index notnull default ''"`
	EventName       string `xorm:"varchar(255) index notnull default ''"`
	Note            string `xorm:"text"`
	LogIndex        int64  `xorm:"integer notnull pk default 0"`
}

type EventHeightCheckpoint struct {
	Id                uint64    `json:"-" xorm:"bigserial pk autoincr"`
	MaxRecordedHeight uint64    `xorm:"bigint notnull"`
	EventHash         string    `xorm:"varchar(255) index notnull"`
	EventName         string    `xorm:"varchar(255)"`
	TxnType           uint8     `xorm:"smallint index comment('0 - transaction, 1 - internal transaction')"`
	UpdatedAt         time.Time `xorm:"updated"`
}

// Receipt evm transaction receipt
type EVMReceipt struct {
	Height            int64  `json:"height"`
	Version           int    `json:"version"`
	TransactionHash   string `json:"transaction_hash"`
	TransactionIndex  int64  `json:"transaction_index"`
	BlockHash         string `json:"block_hash"`
	BlockNumber       int64  `json:"block_number"`
	From              string `json:"from"`
	To                string `json:"to"`
	StateRoot         string `json:"state_root"`
	Status            int64  `json:"status"`
	ContractAddress   string `json:"contract_address"`
	CumulativeGasUsed int64  `json:"cumulative_gas_used"`
	GasUsed           int64  `json:"gas_used"`
	EffectiveGasPrice int64  `json:"effective_gas_price"`
	LogsBloom         string `json:"logs_bloom"`
	Logs              string `json:"logs"`
}

func (r *EVMReceipt) TableName() string {
	return "evm_receipt"
}

// InternalTX contract internal transaction
type EVMInternalTXN struct {
	Height           int64  `json:"height"`
	Version          int    `json:"version"`
	ParentHash       string `json:"parent_hash"`
	ParentMessageCid string `json:"parent_message_cid"`
	Type             uint64 `json:"type"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Params           string `json:"params"`
	ParamsCodec      uint64 `json:"params_codec"`
}

func (r *EVMInternalTXN) TableName() string {
	return "evm_internal_tx"
}
