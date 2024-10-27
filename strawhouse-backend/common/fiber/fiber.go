package fiber

import (
	"context"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"go.uber.org/fx"
)

func Init(lc fx.Lifecycle, config *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler:                 ErrorHandler,
		Prefork:                      false,
		StrictRouting:                true,
		AppName:                      "strawhouse/" + gut.Commit,
		ServerHeader:                 "strawhouse/" + gut.Commit,
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
