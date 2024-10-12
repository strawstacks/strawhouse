package client

import (
	"github.com/spf13/cobra"
	"github.com/strawstacks/strawhouse/strawhouse-command/common"
)

var MetadataDirectoryListCmd = &cobra.Command{
	Use:   "directory-list",
	Short: "List directory",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()
		directory, _ := cmd.Flags().GetString("directory")
		common.Handle(common.Driver.Client.DirectoryList(directory))
	},
}

func init() {
	MetadataDirectoryListCmd.Flags().String("directory", "", "Directory to list")
	_ = MetadataDirectoryListCmd.MarkFlagRequired("directory")
	Cmd.AddCommand(MetadataDirectoryListCmd)
}
