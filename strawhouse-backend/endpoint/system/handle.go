package system

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/name"
	"github.com/strawstacks/strawhouse/strawhouse-driver"
)

type Handler struct {
	Config    *config.Config
	Name      *name.Name
	Signature *strawhouse.Signature
	Pogreb    *pogreb.Pogreb
}

func NewHandler(config *config.Config, pogreb *pogreb.Pogreb, name *name.Name, signature *strawhouse.Signature) *Handler {
	return &Handler{
		Config:    config,
		Pogreb:    pogreb,
		Name:      name,
		Signature: signature,
	}
}
