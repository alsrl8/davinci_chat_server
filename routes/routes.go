package routes

import (
	"davinci-chat/handlers"
	"davinci-chat/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Root)
	app.Get("/ping", handlers.Ping)
	app.Use("/ws", middlewares.JWTMiddleware)
	app.Get("/ws", handlers.Websocket, handlers.Ws)
	app.Post("/login", handlers.LoginHandler)
}
