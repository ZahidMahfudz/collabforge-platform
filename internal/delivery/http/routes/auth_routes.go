package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/internal/delivery/http/controller"
)

func AuthRoutes(router fiber.Router, authController *controller.AuthController) {
	auth := router.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/google", authController.LoginGoogle)
}
