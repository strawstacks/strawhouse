package system

import (
	"backend/common/config"
	"backend/common/pogreb"
	"backend/util/name"
	"backend/util/signature"
)

type Handler struct {
	Config    *config.Config
	Name      *name.Name
	Signature *signature.Signature
	Pogreb    *pogreb.Pogreb
}

func NewHandler(config *config.Config, pogreb *pogreb.Pogreb, name *name.Name, signature *signature.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Name:      name,
		Signature: signature,
	}
}
