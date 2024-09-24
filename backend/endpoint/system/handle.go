package system

import (
	"backend/common/config"
	"backend/util/name"
	"backend/util/signature"
	"github.com/akrylysov/pogreb"
)

type Handler struct {
	Config    *config.Config
	Name      *name.Name
	Signature *signature.Signature
	Pogreb    *pogreb.DB
}

func NewHandler(config *config.Config, pogreb *pogreb.DB, name *name.Name, signature *signature.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Name:      name,
		Signature: signature,
	}
}
