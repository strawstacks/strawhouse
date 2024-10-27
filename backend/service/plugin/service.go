package plugin

import (
	"context"
	"github.com/strawstacks/strawhouse-go"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/service/file"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"go.uber.org/fx"
)

type Store struct {
	Plugins map[string]strawhouse.Plugin
}

type Service struct {
	config    *config.Config
	file      *file.Service
	eventfeed *eventfeed.EventFeed
	s         *Store
}

func Serve(lc fx.Lifecycle, config *config.Config, file *file.Service, eventfeed *eventfeed.EventFeed) *Service {
	service := &Service{
		config:    config,
		file:      file,
		eventfeed: eventfeed,
		s: &Store{
			Plugins: make(map[string]strawhouse.Plugin),
		},
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			service.Init()
			return nil
		},
		OnStop: func(context.Context) error {
			for _, plug := range service.s.Plugins {
				plug.Unload()
			}
			return nil
		},
	})

	return service
}
