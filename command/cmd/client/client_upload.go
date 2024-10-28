package client

import (
	"github.com/bsthun/gut"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strawhouse-command/common"
)

var TransferUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()
		src, _ := cmd.Flags().GetString("src")
		dst, _ := cmd.Flags().GetString("dst")
		name := filepath.Base(dst)
		directory := filepath.Dir(dst)

		// * Read file
		content, err := os.ReadFile(src)
		if err != nil {
			gut.Fatal("unable to open file", err)
		}

		if err := common.Driver.Client.TransferUpload(name, directory, content, nil); err != nil {
			gut.Fatal("unable to upload file", err)
		}

		common.Handle("file uploaded to "+dst, nil)
	},
}

func init() {
	TransferUploadCmd.Flags().String("src", "", "Source file")
	TransferUploadCmd.Flags().String("dst", "", "Destination path (including filename)")
	_ = TransferUploadCmd.MarkFlagRequired("src")
	_ = TransferUploadCmd.MarkFlagRequired("dst")
	Cmd.AddCommand(TransferUploadCmd)
}
