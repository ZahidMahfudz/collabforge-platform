package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/usecase"
	"github.com/zahidmahfudz/collabforge-platform/utils/response"
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
		if err.Error() == "EMAIL_ALREADY_EXISTS" {
			return response.Error(ctx, fiber.StatusConflict, "email sudah ada", "EMAIL_ALREADY_EXISTS")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "internal server error", "INTERNAL_SERVER_ERROR")
	}
	Logger.Debug("usecase register berhasil, menunggu response untuk dikirim ke client")
	

	Logger.Debug("register berhasil, mengirim response")
	return response.Success(ctx, fiber.StatusOK, "register success", result)
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	Logger.Debug("memasuki fungsi Login di AuthController")

	// ambil validated body dari middleware
	req := ctx.Locals("validatedBody").(request.LoginRequest)
	Logger.Debugf("data yang diterima: %+v", req)

	//mengirim data ke usecase untuk proses login
	Logger.Debug("Mengirim data ke usecase untuk proses login")
	result, refreshToken, err := c.authUseCase.Login(ctx.Context(), req)
	if err != nil {
		Logger.Errorf("Error saat login: %v", err)
		if err.Error() == "INVALID_CREDENTIALS" {
			return response.Error(ctx, fiber.StatusUnauthorized, "email atau password salah", "INVALID_CREDENTIALS")
		}
		if err.Error() == "FAILED_TO_GENERATE_TOKEN" {
			return response.Error(ctx, fiber.StatusInternalServerError, "gagal menghasilkan token", "FAILED_TO_GENERATE_TOKEN")
		}

		return response.Error(ctx, fiber.StatusInternalServerError, "internal server error", "INTERNAL_SERVER_ERROR")
	}
	Logger.Debug("usecase login berhasil, menunggu response untuk dikirim ke client")

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: "lax",
	})

	Logger.Debug("login berhasil, mengirim response")
	return response.Success(ctx, fiber.StatusOK, "login success", result)

}