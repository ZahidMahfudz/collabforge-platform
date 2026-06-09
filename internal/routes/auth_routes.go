package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/internal/controller"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/middleware"
)

func AuthRoutes(app *fiber.App, authController *controller.AuthController, authMiddleware *middleware.AuthMiddleware) {
	//route group untuk auth
	authGroup := app.Group("/auth")

	//endpoint register
	authGroup.Post("/register", middleware.ValidateRequest[request.RegisterRequest](), authController.Register)
	//endpoint login
	authGroup.Post("/login", middleware.ValidateRequest[request.LoginRequest](), authController.Login)
	// endpoint refresh token
	authGroup.Post("/refresh", authController.RefreshToken)
	// endpoint logout
	authGroup.Post("/logout", authController.Logout)

	//testing auth middleware, endpoint ini hanya bisa diakses dengan token yang valid
	authGroup.Get("/protected", authMiddleware.Protect(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "you have access to this protected route",
		})
	})

}