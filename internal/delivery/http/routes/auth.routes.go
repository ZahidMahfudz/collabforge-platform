package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Login Page")
	})

	
}