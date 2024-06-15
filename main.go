package main

import (
	"davinci-chat/auth"
	"davinci-chat/middlewares"
	"davinci-chat/routes"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	log.Println("Starting server...")
	auth.InitFirebase()

	app := fiber.New()

	app.Use(middlewares.NewCORS())
	app.Use(middlewares.NewLimiter())

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
