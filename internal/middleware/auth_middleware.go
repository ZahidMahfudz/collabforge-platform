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

func NewAuthMiddleware(pasetoService *token.PasetoService) *AuthMiddleware {
	return &AuthMiddleware{pasetoService: pasetoService}
}

func (m *AuthMiddleware) Protect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		Logger.Debug("memasuki middleware Protect untuk otentikasi request")
		// ambil token dari header Authorization
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			Logger.Debug("header Authorization tidak ditemukan")
			return response.Error(c, fiber.StatusUnauthorized, "missing token", "UNAUTHORIZED")
		}

		//format: "Bearer <token>"
		splitted := strings.Split(authHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			Logger.Debug("format header Authorization tidak valid")
			return response.Error(c, fiber.StatusUnauthorized, "invalid token format", "UNAUTHORIZED")
		}

		tokenStr := splitted[1]

		// validasi token
		claims, err := m.pasetoService.VerifyToken(tokenStr)
		if err != nil {
			Logger.Debugf("validasi token gagal: %v", err)
			return response.Error(c, fiber.StatusUnauthorized, "invalid or expired token", "UNAUTHORIZED")
		}

		// simpan claims ke locals untuk digunakan di handler
		userID, err := claims.GetString("user_id")
		if err != nil {
			Logger.Debugf("gagal mendapatkan user_id dari claims: %v", err)
			return response.Error(c, fiber.StatusUnauthorized, "invalid token", "UNAUTHORIZED")
		}
		
		email, err := claims.GetString("email")
		if err != nil {
			Logger.Debugf("gagal mendapatkan email dari claims: %v", err)
			return response.Error(c, fiber.StatusUnauthorized, "invalid token", "UNAUTHORIZED")
		}

		Logger.Debugf("token valid, user_id: %s, email: %s", userID, email)

		//inject ke context locals
		c.Locals("userID", userID)
		c.Locals("email", email)

		return c.Next()
	}
}

