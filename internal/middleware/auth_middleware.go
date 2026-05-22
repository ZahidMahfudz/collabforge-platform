package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/zahidmahfudz/collabforge-platform/internal/service/token"
	"github.com/zahidmahfudz/collabforge-platform/utils/response"
)

type AuthMiddleware struct {
	pasetoService *token.PasetoService
}

func NewAuthMiddleware(pasetoService *token.PasetoService,) *AuthMiddleware {
	return &AuthMiddleware{pasetoService: pasetoService,}
}

func (m *AuthMiddleware) Protect() fiber.Handler {

	return func(c *fiber.Ctx) error {

		Logger.Debug("memasuki middleware Protect",)

		// ambil authorization header
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			Logger.Debug("Authorization header tidak ditemukan")
			return response.Error(c, fiber.StatusUnauthorized, "missing token", "UNAUTHORIZED")
		}

		// format: Bearer <token>
		splitted := strings.Split(authHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			Logger.Debug("format Authorization tidak valid",)
			return response.Error(c, fiber.StatusUnauthorized, "invalid token format", "UNAUTHORIZED")
		}
		tokenStr := splitted[1]

		// verify token
		claims, err := m.pasetoService.VerifyToken(tokenStr)
		if err != nil {
			Logger.Debugf("token invalid: %v", err)
			return response.Error(c, fiber.StatusUnauthorized, "invalid or expired token", "UNAUTHORIZED")
		}

		// ambil subject
		userID, err := claims.GetSubject()
		if err != nil {
			Logger.Debugf("gagal mendapatkan subject: %v", err)
			return response.Error(c, fiber.StatusUnauthorized, "invalid token", "UNAUTHORIZED")
		}

		// ambil email custom claim
		email, err := claims.GetString("email")
		if err != nil {
			Logger.Debugf("gagal mendapatkan email: %v", err)
			return response.Error(c, fiber.StatusUnauthorized, "invalid token", "UNAUTHORIZED")
		}
		Logger.Debugf("token valid | userID=%s | email=%s", userID, email)

		// inject ke context
		c.Locals("userID", userID)
		c.Locals("email", email)

		return c.Next()
	}
}