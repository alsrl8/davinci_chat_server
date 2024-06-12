package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/websocket/v2"
	"os"
	"time"
)

func main() {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Game Server!")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				return
			}
			if err = c.WriteMessage(mt, msg); err != nil {
				return
			}
		}
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
