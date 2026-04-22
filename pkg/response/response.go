package response

import (
	"github.com/gofiber/fiber/v2"
	AppError "github.com/zahidmahfudz/collabforge-platform/pkg/errors"
)

type APIResponse struct {
	Success bool `json:"success"`
	Massage string `json:"massage"`
	Data interface{} `json:"data,omitempty"`
	Error interface() `json:"error,omitempty"`
}

// success response
func Success(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(APIResponse{
		Success: true,
		Massage: message,
		Data: data,
	})
}

// error response
func Error(c *fiber.Ctx, err error) error {
	//error custom
	if appErr, ok := err.(*AppError.AppError); ok {
		return c.Status(appErr.Status).JSON(APIResponse{
			Success: false,
			Massage: appErr.Message,
			Error: err,
		})
	}
	//fallback error
	return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
		Success: false,
		Massage: "Internal Server Error",
		Error: err,
	})
}
