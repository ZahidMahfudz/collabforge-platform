package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		// frontend yang diizinkan
		AllowOrigins: "http://localhost:3000,http://127.0.0.1:5500",

		// metode HTTP yang diizinkan
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",

		// header yang diizinkan
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",

		// wajib true kalau pakai cookie auth
		AllowCredentials: true,
	})
}

