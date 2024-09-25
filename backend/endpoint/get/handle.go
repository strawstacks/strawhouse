package get

import (
	"github.com/strawstacks/strawhouse/backend/common/config"
	"github.com/strawstacks/strawhouse/backend/common/pogreb"
	"github.com/strawstacks/strawhouse/backend/util/signature"
)

type Handler struct {
	Config    *config.Config
	Pogreb    *pogreb.Pogreb
	Signature *signature.Signature
}

func NewHandler(config *config.Config, pogreb *pogreb.Pogreb, signature *signature.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Signature: signature,
	}
}
