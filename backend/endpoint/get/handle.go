package get

import "backend/common/config"

type Handler struct {
	Config *config.Config
}

func NewHandler(config *config.Config) *Handler {
	return &Handler{
		Config: config,
	}
}
