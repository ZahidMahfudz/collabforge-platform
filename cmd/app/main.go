package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/delivery/http/controller"
	"github.com/zahidmahfudz/collabforge-platform/internal/delivery/http/routes"
	"github.com/zahidmahfudz/collabforge-platform/internal/middleware"
	"github.com/zahidmahfudz/collabforge-platform/internal/repository"
	"github.com/zahidmahfudz/collabforge-platform/internal/usecase"
)

func main() {
	//Inisialisasi config
	config.LoadEnv()    //env
	config.InitLogger() //logger
	db := config.ConnectDB() //database

	// Dependency Injection
	userRepo := repository.NewUserRepository(db)
	secretKey := config.GetEnv("PASETO_SECRET_KEY")
	googleClientID := config.GetEnv("GOOGLE_CLIENT_ID")
	authUsecase := usecase.NewAuthUsecase(userRepo, secretKey, googleClientID)
	authController := controller.NewAuthController(authUsecase)

	//inisialisasi fiber app
	app := fiber.New()

	//logger sederhana untuk setiap request
	app.Use(middleware.RequestLogger())

	//inisialisasi routes
	routes.AuthRoutes(app, authController)

	//jalankan server
	appPort := config.GetEnv("APP_PORT")
	config.Log.Infof("Starting server on port %s", appPort)
	app.Listen(":" + appPort)
}