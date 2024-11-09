package get

import (
	"github.com/strawst/strawhouse-go"
	"strawhouse-backend/common/config"
	"strawhouse-backend/common/pogreb"
	"strawhouse-backend/service/file"
	"strawhouse-backend/util/eventfeed"
	"strawhouse-backend/util/fileflag"
	"strawhouse-backend/util/filepath"
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
