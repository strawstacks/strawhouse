package file

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/fileflag"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/filepath"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
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
