package get

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
)

type Handler struct {
	Config    *config.Config
	Pogreb    *pogreb.Pogreb
	Signature *strawhouse.Signature
}

func NewHandler(config *config.Config, pogreb *pogreb.Pogreb, signature *strawhouse.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Signature: signature,
	}
}
