package routes

import (
	"davinci-chat/handlers"
	"davinci-chat/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middlewares.RedirectHTTPtoHTTPS)
	app.Get("/", handlers.Root)
	app.Get("/ping", handlers.Ping)
	app.Post("/auto-login", handlers.AutoLoginHandler)
	app.Post("/login", handlers.LoginHandler)
	app.Use("/ws", middlewares.JWTMiddleware)
	app.Get("/ws", handlers.Websocket, handlers.Ws)
	app.Post("/new-user", handlers.AddNewUser)
	app.Post("/user-validation", handlers.ValidateUser)
}
