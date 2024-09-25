package main

import (
	"command/cmd/config"
	"command/cmd/sign"
	"command/common"
	"github.com/spf13/cobra"
	"os"
)

var cmd = &cobra.Command{
	Use:   "strawc",
	Short: "Strawhouse CLI for managing config and key operations",
}

func main() {
	common.InitConfig()
	cmd.AddCommand(config.Cmd)
	cmd.AddCommand(sign.Cmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
