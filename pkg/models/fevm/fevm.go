package fevm

type FevmEvent struct {
	Id              uint64 `json:"-" xorm:"bigserial pk autoincr"`
	Height          uint64 `xorm:"bigint notnull"`
	TransactionHash string `xorm:"varchar(255) nootnull pk"`
	From            string `xorm:"varchar(255) index notnull default ''"`
	To              string `xorm:"varchar(255) index notnull default ''"`
	Status          int64  `xorm:"integer notnull default 0"`
	LogsBloom       string `xorm:"text notnull default ''"`
	Logs            string `xorm:"text notnull default ''"`
	EventHash       string `xorm:"varchar(255) index notnull default ''"`
	EventName       string `xorm:"varchar(255) index notnull default ''"`
	Note            string `xorm:"text"`
}
