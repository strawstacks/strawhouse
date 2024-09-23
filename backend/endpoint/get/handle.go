package get

import (
	"backend/common/config"
	"backend/util/signature"
)

type Handler struct {
	Config    *config.Config
	Signature *signature.Signature
}

func NewHandler(config *config.Config, signature *signature.Signature) *Handler {
	return &Handler{
		Config:    config,
		Signature: signature,
	}
}
