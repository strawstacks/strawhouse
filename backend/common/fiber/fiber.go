package fiber

import (
	"context"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"strawhouse-backend/common/config"
)

func Init(lc fx.Lifecycle, config *config.Config) *fiber.App {
	name := "Strawhouse"
	if gut.Version != "" {
		name += " " + gut.Version
	}
	if gut.Commit != "" {
		name += " (" + gut.Commit + ")"
	}
	app := fiber.New(fiber.Config{
		ErrorHandler:                 ErrorHandler,
		Prefork:                      false,
		StrictRouting:                true,
		AppName:                      name,
		ServerHeader:                 name,
		StreamRequestBody:            true,
		DisablePreParseMultipartForm: true,
		Network:                      *config.WebListen[0],
	})

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := app.Listen(*config.WebListen[1])
				if err != nil {
					gut.Fatal("Unable to listen", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			// * Shutdown
			_ = app.Shutdown()
			return nil
		},
	})

	return app
}
