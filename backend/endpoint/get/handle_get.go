package get

import "github.com/gofiber/fiber/v2"

func (r *Handler) Get(c *fiber.Ctx) error {
	return c.JSON("")
}
