package main

import (
	"event-trace/internal/busi"

	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// @title starboard fevm event tracing job
// @version 1.0
// @description starboard event tracing job
// @termsOfService http://swagger.io/terms/

// @contact.name xueyouchen
// @contact.email xueyou@starboardventures.io

// @host localhost:7001
// @BasePath /api/v1
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "event",
		Short: "event tracing",
		Run: func(cmd *cobra.Command, args []string) {
			if err := entry(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringVar(&busi.Flags.Config, "conf", "", "path of the configuration file")

	return cmd
}

func entry() error {
	busi.Start()
	return nil
}

func main() {
	if err := NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
