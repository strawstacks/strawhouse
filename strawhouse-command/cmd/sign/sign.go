package sign

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strawstacks/strawhouse/strawhouse-command/common"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
	"log"
	"strings"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "sign",
	Short: "Generate signed token",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()
		// * Parse action
		actionFlag, _ := cmd.Flags().GetString("action")
		var action strawhouse.SignatureAction
		if actionFlag == "get" {
			action = strawhouse.SignatureActionGet
		} else if actionFlag == "upload" {
			action = strawhouse.SignatureActionUpload
		} else {
			log.Fatalf("action must be one of 'get' or 'upload'")
		}

		// * Parse mode
		modeFlag, _ := cmd.Flags().GetString("mode")
		var mode strawhouse.SignatureMode
		if modeFlag == "file" {
			mode = strawhouse.SignatureModeFile
		} else if modeFlag == "dir" {
			mode = strawhouse.SignatureModeDirectory
		} else {
			log.Fatalf("mode must be one of 'file' or 'dir'")
		}

		// * Parse depth
		recursive, _ := cmd.Flags().GetBool("recursive")

		// * Expired seconds
		expired, _ := cmd.Flags().GetInt("expired")
		if expired <= 0 {
			log.Fatalf("expired must be a positive integer in seconds")
		}

		// * Parse path
		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			log.Fatalf("path is required")
		}
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		if mode == strawhouse.SignatureModeDirectory && !strings.HasSuffix(path, "/") {
			path += "/"
		}

		// * Generate signed token
		token := common.Driver.Signature.Generate(action, mode, path, recursive, time.Now().Add(time.Duration(expired)*time.Second), nil)
		fmt.Println(token)
	},
}

func init() {
	Cmd.Flags().String("action", "", "Action to perform. One of 'get' or 'upload'")
	Cmd.Flags().String("mode", "", "Entity mode. One of 'file' or 'dir'")
	Cmd.Flags().Bool("recursive", false, "Allow recursive access")
	Cmd.Flags().Int("expired", 0, "Expired seconds")
	Cmd.Flags().String("path", "", "Path to sign")

	_ = Cmd.MarkFlagRequired("action")
	_ = Cmd.MarkFlagRequired("mode")
	_ = Cmd.MarkFlagRequired("recursive")
	_ = Cmd.MarkFlagRequired("expired")
	_ = Cmd.MarkFlagRequired("path")
}
