package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/utils"
)


func ValidateRequest[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		Logger.Debugf("memasuki fungsi ValidateRequest dengan parameter tipe: %T", *new(T))

		var body T

		Logger.Debugf("mulai parse request body untuk tipe: %T", *new(T))
		// parse request body
		if err := c.BodyParser(&body); err != nil {
			Logger.Debugf("gagal parse request body: %v", err)

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "invalid request body",
				"error":   "BAD_REQUEST",
			})
		}

		Logger.Debugf("sukses parse request body untuk tipe: %T", *new(T))
		Logger.Debugf("mulai validasi struct untuk tipe: %T", *new(T))
		// validate struct
		if err := utils.Validate.Struct(body); err != nil {
			Logger.Debugf("validasi gagal: %v", err)

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "validation failed",
				"errors":  utils.FormatValidationError(err),
			})
		}

		Logger.Debugf("validasi sukses untuk tipe: %T", *new(T))
		Logger.Debugf("menyimpan body tervalidasi ke locals")
		// simpan body ke locals
		c.Locals("validatedBody", body)

		return c.Next()
	}
}