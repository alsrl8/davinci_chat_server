package middlewares

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewCORS() fiber.Handler {
	env := config.GetRunEnv()
	allowOrigin := getAllowOrigin(env)
	return cors.New(cors.Config{
		AllowOrigins:     allowOrigin,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	})
}

func getAllowOrigin(env consts.RunEnv) (allowOrigin string) {
	switch env {
	case consts.Production:
		allowOrigin = "https://songmingi.com"
	case consts.Development:
		allowOrigin = "http//localhost:3000"
	}
	return
}
