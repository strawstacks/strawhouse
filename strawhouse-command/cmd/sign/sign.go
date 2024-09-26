package sign

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strawstacks/strawhouse/strawhouse-command/common"
	"log"
	"strings"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "sign",
	Short: "Generate signed token",
	Run: func(cmd *cobra.Command, args []string) {
		common.InitDriver()

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

		// * Parse depth
		depth, _ := cmd.Flags().GetInt("depth")
		if depth < 0 {
			log.Fatalf("depth must be a positive integer")
		}

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
		if !strings.HasSuffix(path, "/") {
			path += "/"
		}

		// * Generate signed token
		token := common.Driver.Signature.Generate(1, mode, action, uint32(depth), time.Now().Add(time.Duration(expired)*time.Second), path, nil)
		fmt.Println(token)
	},
}

func init() {
	Cmd.Flags().String("mode", "", "Entity mode. One of 'file' or 'dir'")
	Cmd.Flags().String("action", "", "Action to perform. One of 'get' or 'upload'")
	Cmd.Flags().Int("depth", 0, "Depth of directory")
	Cmd.Flags().Int("expired", 0, "Expired seconds")
	Cmd.Flags().String("path", "", "Path to sign")

	_ = Cmd.MarkFlagRequired("mode")
	_ = Cmd.MarkFlagRequired("action")
	_ = Cmd.MarkFlagRequired("depth")
	_ = Cmd.MarkFlagRequired("expired")
	_ = Cmd.MarkFlagRequired("path")
}
