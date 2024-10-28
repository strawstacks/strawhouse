package main

import (
	"github.com/spf13/cobra"
	"os"
	"strawhouse-command/cmd/client"
	"strawhouse-command/cmd/config"
	"strawhouse-command/cmd/sign"
	"strawhouse-command/common"
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
