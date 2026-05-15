package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/internal/controller"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/middleware"
)

func AuthRoutes(app *fiber.App, authController *controller.AuthController) {
	//route group untuk auth
	authGroup := app.Group("/auth")

	//endpoint register
	authGroup.Post("/register", middleware.ValidateRequest[request.RegisterRequest](), authController.Register)


}