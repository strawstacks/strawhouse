package client

import (
	"github.com/spf13/cobra"
	"github.com/strawstacks/strawhouse/strawhouse-command/common"
	"github.com/strawstacks/strawhouse/strawhouse-proto/pb"
)

var FeedUploadCmd = &cobra.Command{
	Use:   "feed-upload",
	Short: "Feed upload events",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()
		directory, _ := cmd.Flags().GetString("directory")
		closeCh := make(chan struct{})
		upload, err := common.Driver.Client.FeedUpload(directory, func(resp *pb.UploadFeedResponse, err error) {
			if err != nil {
				common.Handle(nil, err)
				close(closeCh)
				return
			}
			common.Handle(resp, nil)
		})
		if err != nil {
			common.Handle(nil, err)
			return
		}
		defer upload.Close()
		select {
		case <-closeCh:
		}
	},
}

func init() {
	FeedUploadCmd.Flags().String("directory", "", "Directory to watch")
	_ = FeedUploadCmd.MarkFlagRequired("directory")
	Cmd.AddCommand(FeedUploadCmd)
}
