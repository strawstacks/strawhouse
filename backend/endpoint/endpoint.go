package endpoint

import (
	"backend/endpoint/get"
	"backend/endpoint/system"
	"github.com/gofiber/fiber/v2"
)

func Bind(app *fiber.App, systemHandler *system.Handler, getHandler *get.Handler) {
	// * System route
	systemRoute := app.Group("_/")
	systemRoute.Get("upload", systemHandler.Upload)

	// * Get route
	app.Get("/*", getHandler.Get)
}
