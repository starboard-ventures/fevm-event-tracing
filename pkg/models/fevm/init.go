package fevm

var (
	Tables []interface{}
)

func init() {
	Tables = append(Tables,
		new(FevmEvent),
		new(EventHeightCheckpoint),
	)
}
