package get

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/fileflag"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
)

type Handler struct {
	Config    *config.Config
	Pogreb    *pogreb.Pogreb
	Fileflag  *fileflag.Fileflag
	EventFeed *eventfeed.EventFeed
	Signature *strawhouse.Signature
}

func NewHandler(config *config.Config, pogreb *pogreb.Pogreb, fileflag *fileflag.Fileflag, signature *strawhouse.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Fileflag:  fileflag,
		Signature: signature,
	}
}
