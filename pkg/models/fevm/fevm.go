package fevm

type FevmEvent struct {
	Id              uint64 `json:"-" xorm:"bigserial pk autoincr"`
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
}

type EventHeightCheckpoint struct {
	Id                uint64 `json:"-" xorm:"bigserial pk autoincr"`
	MaxRecordedHeight uint64 `xorm:"bigint notnull"`
	EventHash         string `xorm:"varchar(255) index notnull"`
	EventName         string `xorm:"varchar(255)"`
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
