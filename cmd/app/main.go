package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/controller"
	"github.com/zahidmahfudz/collabforge-platform/internal/middleware"
	"github.com/zahidmahfudz/collabforge-platform/internal/repository"
	"github.com/zahidmahfudz/collabforge-platform/internal/routes"
	"github.com/zahidmahfudz/collabforge-platform/internal/service/token"
	"github.com/zahidmahfudz/collabforge-platform/internal/usecase"
	"github.com/zahidmahfudz/collabforge-platform/utils"
)

func main() {
	//Inisialisasi config
	config.LoadEnv()    //env
	config.InitLogger() //logger
	db := config.ConnectDB()  //database
	utils.InitValidator() //validator

	//dependency injection untuk repository, usecase, dan controller

	//inisialisasi repository, usecase, dan controller untuk auth
	userRepo := repository.NewUserRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	pasetoService := token.NewPasetoService()
	authUseCase := usecase.NewAuthUseCase(userRepo, refreshTokenRepo, pasetoService)
	authController := controller.NewAuthController(authUseCase)
	authMiddleware := middleware.NewAuthMiddleware(pasetoService)

	//inisialisasi repository, usecase, dan controller untuk fitur lain bisa ditambahkan di sini dengan pola yang sama

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

	//daftarkan route untuk auth
	routes.AuthRoutes(app, authController, authMiddleware)

	//jalankan server
	appPort := config.GetEnv("APP_PORT")
	Logger.Infof("Starting server on port %s", appPort)
	err := app.Listen(":" + appPort)
	if err != nil {
		Logger.Fatal("Failed to start server")
	}
}