package main

import (
	"github.com/spf13/cobra"
	"github.com/strawstacks/strawhouse/strawhouse-command/cmd/client"
	"github.com/strawstacks/strawhouse/strawhouse-command/cmd/config"
	"github.com/strawstacks/strawhouse/strawhouse-command/cmd/sign"
	"github.com/strawstacks/strawhouse/strawhouse-command/common"
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
	cmd.AddCommand(client.Cmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
