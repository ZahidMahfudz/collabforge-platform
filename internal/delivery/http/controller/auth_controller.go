package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/usecase"
	"github.com/zahidmahfudz/collabforge-platform/pkg/response"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) *AuthController {
	return &AuthController{authUsecase: authUsecase}
}

func (h *AuthController) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, err)
	}

	res, err := h.authUsecase.Register(c.Context(), req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, fiber.StatusCreated, "register success", res)
}

func (h *AuthController) LoginGoogle(c *fiber.Ctx) error {
	var req request.GoogleLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, err)
	}

	res, err := h.authUsecase.LoginGoogle(c.Context(), req)
	if err != nil {
		return response.Error(c, err)
	}

	return response.Success(c, fiber.StatusOK, "login success", res)
}
