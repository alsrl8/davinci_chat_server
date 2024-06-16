package main

import (
	"davinci-chat/auth"
	"davinci-chat/middlewares"
	"davinci-chat/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	log.Println("Starting server...(log)")
	fmt.Println("Starting server...(stdout)")
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
