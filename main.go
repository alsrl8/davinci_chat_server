package main

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"davinci-chat/database"
	"davinci-chat/logx"
	"davinci-chat/middlewares"
	"davinci-chat/routes"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	env := config.GetRunEnv()

	logger := logx.GetLogger()
	defer logger.Close()
	logger.Info("Running on ", env)

	app := fiber.New()

	app.Use(middlewares.NewCORS())
	app.Use(middlewares.NewLimiter())

	routes.SetupRoutes(app)

	db := database.GetDatabase()
	defer func(db database.Database) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)

	switch env {
	case consts.Production:
		certPath := os.Getenv("CERT_PATH")
		keyPath := os.Getenv("KEY_PATH")
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		logger.Fatal(app.ListenTLS(":"+port, certPath, keyPath))
	case consts.Development:
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		logger.Fatal(app.Listen("localhost:" + port))
	}
}
