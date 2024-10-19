package client

import (
	"github.com/bsthun/gut"
	"github.com/spf13/cobra"
	"github.com/strawstacks/strawhouse/strawhouse-command/common"
	"os"
	"path/filepath"
)

var TransferUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()
		src := cmd.Flag("src").Value.String()
		dest, _ := cmd.Flags().GetString("dest")
		name := filepath.Base(dest)
		directory := filepath.Dir(dest)

		// * Read file
		content, err := os.ReadFile(src)
		if err != nil {
			gut.Fatal("unable to open file", err)
		}

		if err := common.Driver.Client.TransferUpload(name, directory, content, nil); err != nil {
			gut.Fatal("unable to upload file", err)
		}

		common.Handle("file uploaded to "+dest, nil)
	},
}

func init() {
	FeedUploadCmd.Flags().String("src", "", "Source file")
	FeedUploadCmd.Flags().String("dest", "", "Destination file")
	_ = FeedUploadCmd.MarkFlagRequired("src")
	_ = FeedUploadCmd.MarkFlagRequired("dest")
	Cmd.AddCommand(TransferUploadCmd)
}
