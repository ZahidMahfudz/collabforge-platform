package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/middleware"
	"github.com/zahidmahfudz/collabforge-platform/routes"
)

func main() {
	//Inisialisasi config
	config.LoadEnv() //env
	config.InitLogger() //logger

	//inisialisasi fiber app
	app := fiber.New()

	//logger sederhana untuk setiap request
	app.Use(middleware.RequestLogger())

	//inisialisasi routes
	routes.AuthRoutes(app)

	//jalankan server
	port := config.GetEnv("APP_PORT")
	app.Listen(":" + port)
	config.Log.Infof("Starting server on port %s", port)
}