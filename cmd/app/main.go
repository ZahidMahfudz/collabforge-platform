package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/delivery/http/routes"
	"github.com/zahidmahfudz/collabforge-platform/internal/middleware"
)

func main() {
	//Inisialisasi config
	config.LoadEnv() //env
	config.InitLogger() //logger
	config.ConnectDB() //database

	//inisialisasi fiber app
	app := fiber.New()

	//logger sederhana untuk setiap request
	app.Use(middleware.RequestLogger())

	//inisialisasi routes
	routes.AuthRoutes(app)

	//jalankan server
	appPort := config.GetEnv("APP_PORT")
	app.Listen(":" + appPort)
	config.Log.Infof("Starting server on port %s", appPort)
}