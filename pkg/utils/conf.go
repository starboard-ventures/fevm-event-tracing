package utils

import (
	"syscall"

	"github.com/jinzhu/configor"
)

var CNF TomlConfig

type TomlConfig struct {
	Title     string
	FevmEvent FevmEvent `toml:"fevm_event"`
}

type FevmEvent struct {
	Addr                string `toml:"listen"`
	DB                  string `toml:"database"`
	Lotus               string `toml:"lotus"`
	WfilContract        string `toml:"wfil_contract"`
	PfilContract        string `toml:"pfil_contract"`
	ReplContract        string `toml:"repl_contract"`
	ReplAuctionContract string `toml:"repl_auction_contract"`
}

func InitConfFile(file string, cf *TomlConfig) error {
	err := syscall.Access(file, syscall.O_RDONLY)
	if err != nil {
		return err
	}
	err = configor.Load(cf, file)
	if err != nil {
		return err
	}

	return nil
}
