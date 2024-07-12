package handlers

import (
	"davinci-chat/types"
	"davinci-chat/utils"
	"encoding/json"
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
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		userName, err := utils.GetUserName(c)
		if err != nil {
			messageObj := types.Message{User: "system", Message: "can't find user information"}
			jsonData, err := json.Marshal(messageObj)
			if err != nil {
				log.Println("json marshal error:", err)
				break
			}
			err = c.WriteMessage(mt, jsonData)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}

		messageObj := types.Message{User: userName, Message: string(msg)}
		jsonData, err := json.Marshal(messageObj)
		if err != nil {
			log.Println("json marshal error:", err)
			break
		}

		err = c.WriteMessage(mt, jsonData)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
})
