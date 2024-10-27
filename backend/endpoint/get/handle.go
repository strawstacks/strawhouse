package get

import (
	"github.com/strawstacks/strawhouse-go"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/service/file"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/fileflag"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/filepath"
)

type Handler struct {
	Config    *config.Config
	Pogreb    *pogreb.Pogreb
	File      *file.Service
	Filepath  *filepath.Filepath
	Fileflag  *fileflag.Fileflag
	EventFeed *eventfeed.EventFeed
	Signature *strawhouse.Signature
}

func NewHandler(config *config.Config, pogreb *pogreb.Pogreb, file *file.Service, filepath *filepath.Filepath, fileflag *fileflag.Fileflag, eventfeed *eventfeed.EventFeed, signature *strawhouse.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		File:      file,
		Filepath:  filepath,
		Fileflag:  fileflag,
		EventFeed: eventfeed,
		Signature: signature,
	}
}
