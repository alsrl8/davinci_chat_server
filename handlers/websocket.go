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
var connUserMap = make(map[*websocket.Conn]types.User)
var userEmailConnMap = make(map[string]*websocket.Conn)

var Ws = websocket.New(func(c *websocket.Conn) {
	logger := logx.GetLogger()

	mutex.Lock()
	if err := connectWebSocket(c); err != nil {
		logger.Info("failed to connect with websocket protocols with user")
		return
	}
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(connections, c)
		userEmail := connUserMap[c].UserEmail
		delete(userEmailConnMap, userEmail)
		delete(connUserMap, c)
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
		messageObj := types.Message{User: userName, Message: string(msg)}
		broadcastMessage(mt, messageObj, c)
	}
})

func connectWebSocket(c *websocket.Conn) error {
	userName, err := utils.GetUserName(c)
	if err != nil {
		return err
	}
	userEmail, err := utils.GetUserEmail(c)
	if err != nil {
		return err
	}
	connections[c] = true
	connUserMap[c] = types.User{
		UserName:  userName,
		UserEmail: userEmail,
	}
	userEmailConnMap[userEmail] = c

	return nil
}

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

		err = conn.WriteMessage(mt, jsonData)
		if err != nil {
			logger.Info("write error on broadcast to %v: %v", conn.RemoteAddr(), err)
			continue
		}
	}
}
