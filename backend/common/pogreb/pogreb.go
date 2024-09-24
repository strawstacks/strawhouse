package pogreb

import (
	"backend/common/config"
	"github.com/akrylysov/pogreb"
	uu "github.com/bsthun/goutils"
	"time"
)

func Init(config *config.Config) *pogreb.DB {
	db, err := pogreb.Open(*config.PogrebPath, &pogreb.Options{
		BackgroundSyncInterval:       1 * time.Minute,
		BackgroundCompactionInterval: 24 * time.Hour,
		FileSystem:                   nil,
	})
	if err != nil {
		uu.Fatal("pogreb error", err)
	}
	return db
}
