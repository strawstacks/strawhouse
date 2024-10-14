package get

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/fileflag"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/filepath"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
)

type Handler struct {
	Config    *config.Config
	Pogreb    *pogreb.Pogreb
	Filepath  *filepath.Filepath
	Fileflag  *fileflag.Fileflag
	EventFeed *eventfeed.EventFeed
	Signature *strawhouse.Signature
}

func NewHandler(config *config.Config, pogreb *pogreb.Pogreb, filepath *filepath.Filepath, fileflag *fileflag.Fileflag, signature *strawhouse.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Filepath:  filepath,
		Fileflag:  fileflag,
		Signature: signature,
	}
}
