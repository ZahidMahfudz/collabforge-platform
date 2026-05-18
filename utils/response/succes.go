package response

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Success  bool      `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *fiber.Ctx,statusCode int,message string, data interface{}) error {
	return ctx.Status(statusCode).JSON(
		SuccessResponse{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}