package main

import (
	"davinci-chat/auth"
	"davinci-chat/config"
	"davinci-chat/middlewares"
	"davinci-chat/routes"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	env := config.GetRunEnv()
	log.Printf("Running on %s\n", env)

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
