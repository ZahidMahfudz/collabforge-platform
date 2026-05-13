package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/middleware"
)

func main() {
	//Inisialisasi config
	config.LoadEnv()    //env
	config.InitLogger() //logger
	config.ConnectDB()  //database

	//simpan variabel logger untuk digunakan dengan mudah
	var Logger = config.Logger

	//inisialisasi fiber app
	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 30,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 30,
	})

	// gunakan middleware CORS untuk mengizinkan akses dari frontend
	app.Use(middleware.CORSMiddleware())

	//logger sederhana untuk setiap request
	app.Use(middleware.RequestLogger())

	app.Get("/", func(ctx *fiber.Ctx) error {
		Logger.Debug("memasuki endpoint root")
		Logger.Debug("mengirim response welcome message")
		return ctx.SendString("Hello, welcome to our API!")
	})

	app.Get("/health", func(ctx *fiber.Ctx) error {
		Logger.Debug("memasuki endpoint health")
		Logger.Debug("mengirim response health check")
		return ctx.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	//jalankan server
	appPort := config.GetEnv("APP_PORT")
	Logger.Infof("Starting server on port %s", appPort)
	err := app.Listen(":" + appPort)
	if err != nil {
		Logger.Fatal("Failed to start server")
	}
}