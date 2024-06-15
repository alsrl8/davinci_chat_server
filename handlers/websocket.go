package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func Websocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUpgradeRequired)
}

var Ws = websocket.New(func(c *websocket.Conn) {
	c.WriteMessage(websocket.TextMessage, []byte("Welcome to the WebSocket server"))

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
})
