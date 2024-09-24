package system

import (
	"backend/common/config"
	"backend/util/signature"
	"github.com/akrylysov/pogreb"
)

type Handler struct {
	Config    *config.Config
	Signature *signature.Signature
	Pogreb    *pogreb.DB
}

func NewHandler(config *config.Config, pogreb *pogreb.DB, signature *signature.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Signature: signature,
	}
}
