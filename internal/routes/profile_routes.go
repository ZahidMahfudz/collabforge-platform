package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/internal/controller"
	"github.com/zahidmahfudz/collabforge-platform/internal/middleware"
)

func ProfileRoutes(app *fiber.App, profileController *controller.ProfileController, authMiddleware *middleware.AuthMiddleware) {
	//route group untuk profile, hanya bisa diakses dengan token yang valid
	profileGroup := app.Group("/profile", authMiddleware.Protect())

	//endpoint untuk mendapatkan URL avatar
	profileGroup.Get("/avatar", profileController.GetAvatarURL)
}
