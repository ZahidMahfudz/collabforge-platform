package response

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func Error(ctx *fiber.Ctx, statusCode int, message string, details interface{}) error {
	return ctx.Status(statusCode).JSON(
		ErrorResponse{
			Success: false,
			Message: message,
			Details: details,
		},
	)
}