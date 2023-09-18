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
	Addr  string `toml:"listen"`
	DB    string `toml:"database"`
	Lotus string `toml:"lotus"`
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
