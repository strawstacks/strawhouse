package file

import (
	"github.com/strawst/strawhouse-go"
	"strawhouse-backend/common/config"
	"strawhouse-backend/common/pogreb"
	"strawhouse-backend/util/eventfeed"
	"strawhouse-backend/util/fileflag"
	"strawhouse-backend/util/filepath"
)

type Service struct {
	config    *config.Config
	pogreb    *pogreb.Pogreb
	fileflag  *fileflag.Fileflag
	filepath  *filepath.Filepath
	eventfeed *eventfeed.EventFeed
	signature *strawhouse.Signature
}

func Serve(config *config.Config, pogreb *pogreb.Pogreb, fileflag *fileflag.Fileflag, filepath *filepath.Filepath, eventfeed *eventfeed.EventFeed, signature *strawhouse.Signature) *Service {
	return &Service{
		config:    config,
		pogreb:    pogreb,
		fileflag:  fileflag,
		filepath:  filepath,
		eventfeed: eventfeed,
		signature: signature,
	}
}
