package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/internal/usecase"
	"github.com/zahidmahfudz/collabforge-platform/utils/response"
)

type ProfileController struct {
	profileUseCase *usecase.ProfileUseCase
}

func NewProfileController(profileUseCase *usecase.ProfileUseCase) *ProfileController {
	return &ProfileController{profileUseCase: profileUseCase}
}

func (c *ProfileController) GetAvatarURL(ctx *fiber.Ctx) error {
	Logger.Debug("memasuki fungsi GetAvatarURL di ProfileController")

	//mengirim data ke usecase untuk mendapatkan URL avatar
	Logger.Debug("Mengirim data ke usecase untuk mendapatkan URL avatar")
	url, err := c.profileUseCase.GetAvatarURL(ctx.Context())
	if err != nil {
		Logger.Errorf("Error saat mendapatkan URL avatar: %v", err)
		return response.Error(ctx, fiber.StatusInternalServerError, "internal server error", "INTERNAL_SERVER_ERROR")
	}

	Logger.Debug("URL avatar berhasil didapatkan, mengirim response")
	return response.Success(ctx, fiber.StatusOK, "avatar URL retrieved successfully", url)
}
