package pogreb

import (
	"context"
	"github.com/akrylysov/pogreb"
	"github.com/bsthun/gut"
	"go.uber.org/fx"
	"path/filepath"
	"strawhouse-backend/common/config"
	"time"
)

type Pogreb struct {
	Sum *pogreb.DB
	Log *pogreb.DB
}

func Init(lc fx.Lifecycle, config *config.Config) *Pogreb {
	options := &pogreb.Options{
		BackgroundSyncInterval:       1 * time.Minute,
		BackgroundCompactionInterval: 24 * time.Hour,
		FileSystem:                   nil,
	}

	sumDb, err := pogreb.Open(filepath.Join(*config.PogrebPath, "sum"), options)
	if err != nil {
		gut.Fatal("pogreb sumdb error", err)
	}

	logDb, err := pogreb.Open(filepath.Join(*config.PogrebPath, "log"), options)
	if err != nil {
		gut.Fatal("pogreb logdb error", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			_ = sumDb.Close()
			_ = logDb.Close()
			return nil
		},
	})

	return &Pogreb{
		Sum: sumDb,
		Log: logDb,
	}
}
