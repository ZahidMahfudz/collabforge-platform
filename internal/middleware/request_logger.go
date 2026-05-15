package middleware

import (
	"time"

	"github.com/zahidmahfudz/collabforge-platform/config"

	"github.com/gofiber/fiber/v2"
)

var Logger = config.Logger

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		Logger.WithFields(map[string]interface{}{
			"method":  c.Method(),
			"path":    c.Path(),
			"status":  c.Response().StatusCode(),
			"latency": time.Since(start).String(),
			"ip":      c.IP(),
		}).Info("transaksi request telah selesai")

		return err
	}
}