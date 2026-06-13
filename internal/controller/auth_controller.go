package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/service"
	"github.com/zahidmahfudz/collabforge-platform/internal/usecase"
	"github.com/zahidmahfudz/collabforge-platform/utils/response"
)

var Logger = config.Logger

type AuthController struct {
	authUseCase   *usecase.AuthUseCase
	googleService *service.GoogleAuthService
}

func NewAuthController(authUseCase *usecase.AuthUseCase, googleService *service.GoogleAuthService) *AuthController {
	return &AuthController{authUseCase: authUseCase, googleService: googleService}
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
	return response.Success(ctx, fiber.StatusCreated, "register success", result)
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

func (c *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	Logger.Debug("memasuki fungsi RefreshToken di AuthController")

	// ambil refresh token dari cookie
	refreshToken := ctx.Cookies("refresh_token")
	if refreshToken == "" {
		Logger.Debug("refresh token tidak ditemukan di cookie")
		return response.Error(ctx, fiber.StatusUnauthorized, "refresh token tidak ditemukan", "REFRESH_TOKEN_NOT_FOUND")
	}
	Logger.Debugf("refresh token ditemukan: %s", refreshToken)

	// kirim ke usecase untuk proses refresh token
	result, newRefreshToken, err := c.authUseCase.RefreshToken(ctx.Context(), refreshToken)
	if err != nil {
		Logger.Errorf("Error saat refresh token: %v", err)
		return response.Error(ctx, fiber.StatusUnauthorized, "refresh token tidak valid", "INVALID_REFRESH_TOKEN")
	}

	// set cookie untuk refresh token baru
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    newRefreshToken,
		HTTPOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: "lax",
	})

	Logger.Debug("refresh token berhasil, mengirim response")
	return response.Success(ctx, fiber.StatusOK, "refresh token success", result)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	Logger.Debug("memasuki fungsi Logout di AuthController")

	// ambil refresh token dari cookie
	refreshToken := ctx.Cookies("refresh_token")
	if refreshToken == "" {
		Logger.Debug("refresh token tidak ditemukan di cookie")
		return response.Error(ctx, fiber.StatusUnauthorized, "refresh token tidak ditemukan", "REFRESH_TOKEN_NOT_FOUND")
	}
	Logger.Debugf("refresh token ditemukan: %s", refreshToken)

	// kirim ke usecase untuk proses logout
	result, err := c.authUseCase.Logout(ctx.Context(), refreshToken)
	if err != nil {
		Logger.Errorf("Error saat logout: %v", err)
		if err.Error() == "INVALID_REFRESH_TOKEN" {
			return response.Error(ctx, fiber.StatusUnauthorized, "refresh token tidak valid", "INVALID_REFRESH_TOKEN")
		}
		if err.Error() == "REFRESH_TOKEN_ALREADY_REVOKED" {
			return response.Error(ctx, fiber.StatusBadRequest, "refresh token sudah direvoke", "REFRESH_TOKEN_ALREADY_REVOKED")
		}
		if err.Error() == "FAILED_TO_REVOKE_TOKEN" {
			return response.Error(ctx, fiber.StatusInternalServerError, "gagal mencabut token", "FAILED_TO_REVOKE_TOKEN")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "internal server error", "INTERNAL_SERVER_ERROR")
	}

	// hapus cookie refresh_token
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: "lax",
		MaxAge:   -1, // menghapus cookie
	})

	Logger.Debug("logout berhasil, cookie dihapus, mengirim response")
	return response.Success(ctx, fiber.StatusOK, "logout success", result)
}

func (c *AuthController) GoogleLogin(ctx *fiber.Ctx) error {
	url := c.googleService.GetLoginURL()

	return ctx.Redirect(url)
}

func (c *AuthController) GoogleCallBack(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	if code == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Code google call back not found", "INTERNAL_SERVER_ERROR")
	}

	user, err := c.googleService.GetuserByCode(code)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Gagal memuat user", "INTERNAL_SERVER_ERROR")
	}

	return response.Success(ctx, fiber.StatusOK, "Google call back success", user)
}
