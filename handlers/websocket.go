package handlers

import (
	"davinci-chat/logx"
	"davinci-chat/types"
	"davinci-chat/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"sync"
)

func Websocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUpgradeRequired)
}

var mutex sync.Mutex

var connections = make(map[*websocket.Conn]bool)

var Ws = websocket.New(func(c *websocket.Conn) {
	logger := logx.GetLogger()

	mutex.Lock()
	connections[c] = true
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(connections, c)
		mutex.Unlock()
		err := c.Close()
		if err != nil {
			logger.Info("close connection error: %v", err)
		}
	}()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			logger.Info("read error: %v", err)
			break
		}

		userName, err := utils.GetUserName(c)
		if err != nil {
			messageObj := types.Message{User: "system", Message: "can't find user information"}
			jsonData, err := json.Marshal(messageObj)
			if err != nil {
				logger.Info("json marshal error: %v", err)
				break
			}
			err = c.WriteMessage(mt, jsonData)
			if err != nil {
				logger.Info("write error: %v", err)
				break
			}
		}

		messageObj := types.Message{User: userName, Message: string(msg)}
		broadcastMessage(mt, messageObj, c)
	}
})

func broadcastMessage(mt int, message types.Message, sender *websocket.Conn) {
	logger := logx.GetLogger()

	jsonData, err := json.Marshal(message)
	if err != nil {
		logger.Info("write error: %v", err)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for conn := range connections {
		if conn == sender {
			continue
		}

		err = conn.WriteMessage(mt, jsonData)
		if err != nil {
			logger.Info("write error on broadcast to %v: %v", conn.RemoteAddr(), err)
			continue
		}
	}
}
