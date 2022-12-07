package server

import (
	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "API is running",
		})
	})

	return app
}
