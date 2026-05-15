package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/usecase"
)

var Logger = config.Logger

type AuthController struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthController(authUseCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{authUseCase: authUseCase}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	Logger.Debug("memasuki fungsi Register di AuthController")
	// ambil validated body dari middleware
	req := ctx.Locals("validatedBody").(request.RegisterRequest)

	Logger.Debugf("data yang diterima: %+v", req)

	//mengirim data ke usecase untuk proses register
	Logger.Debug("Mengirim data ke usecase untuk proses register")
	result, err := c.authUseCase.Register(ctx.Context(), req)
	if err != nil {
		Logger.Errorf("Error saat register: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "register failed",
			"error":   err.Error(),
		})
	}
	Logger.Debug("usecase register berhasil, menunggu response untuk dikirim ke client")
	

	Logger.Debug("register berhasil, mengirim response")
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "register success",
		"data": result,
	})
}