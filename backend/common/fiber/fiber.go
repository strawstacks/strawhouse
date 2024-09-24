package fiber

import (
	"backend/common/config"
	"context"
	uu "github.com/bsthun/goutils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func Init(lc fx.Lifecycle, config *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler:      ErrorHandler,
		Prefork:           false,
		StrictRouting:     true,
		StreamRequestBody: true,
		Network:           *config.WebListen[0],
	})

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := app.Listen(*config.WebListen[1])
				if err != nil {
					uu.Fatal("Unable to listen", err)
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
