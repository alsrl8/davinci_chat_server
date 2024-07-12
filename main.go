package main

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"davinci-chat/database"
	"davinci-chat/middlewares"
	"davinci-chat/routes"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {

	env := config.GetRunEnv()
	log.Printf("Running on %s\n", env)

	app := fiber.New()

	//app.Use(middlewares.NewCORS())
	app.Use(middlewares.NewLimiter())

	routes.SetupRoutes(app)

	db := database.GetDatabase()
	defer func(db database.Database) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	switch env {
	case consts.Production:
		certPath := os.Getenv("CERT_PATH")
		keyPath := os.Getenv("KEY_PATH")
		log.Fatal(app.ListenTLS("chat.songmingi.com:8080", certPath, keyPath))
	case consts.Development:
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Fatal(app.Listen("localhost:" + port))
	}
}
