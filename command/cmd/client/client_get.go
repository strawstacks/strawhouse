package client

import (
	"github.com/bsthun/gut"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strawhouse-command/common"
)

var TransferGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a file",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()
		path, _ := cmd.Flags().GetString("path")

		resp, err := common.Driver.Client.TransferGet(path)
		if err != nil {
			gut.Fatal("unable to get file", err)
		}

		name := filepath.Base(path)
		if err := os.WriteFile(name, resp.Content, 0644); err != nil {
			gut.Fatal("unable to write file", err)
		}

		common.Handle("file downloaded to "+name, nil)
	},
}

func init() {
	TransferGetCmd.Flags().String("path", "", "Remote file path")
	_ = TransferGetCmd.MarkFlagRequired("path")
	Cmd.AddCommand(TransferGetCmd)
}
