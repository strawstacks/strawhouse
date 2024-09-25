package get

import (
	"backend/common/config"
	"backend/common/pogreb"
	"backend/util/signature"
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
