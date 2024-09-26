package pogreb

import (
	"context"
	"github.com/akrylysov/pogreb"
	uu "github.com/bsthun/goutils"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"go.uber.org/fx"
	"path/filepath"
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
		uu.Fatal("pogreb sumdb error", err)
	}

	logDb, err := pogreb.Open(filepath.Join(*config.PogrebPath, "log"), options)
	if err != nil {
		uu.Fatal("pogreb logdb error", err)
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
