package client

import (
	"github.com/spf13/cobra"
	"github.com/strawstacks/strawhouse/command/common"
)

var Cmd = &cobra.Command{
	Use:   "client",
	Short: "Client command",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()
	},
}

func init() {
}
