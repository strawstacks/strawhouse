package system

import (
	"github.com/strawstacks/strawhouse/backend/common/config"
	"github.com/strawstacks/strawhouse/backend/common/pogreb"
	"github.com/strawstacks/strawhouse/backend/util/name"
	"github.com/strawstacks/strawhouse/backend/util/signature"
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
