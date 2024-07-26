package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CountActiveChatUser(c *fiber.Ctx) error {
	mutex.Lock()
	number := len(connections)
	mutex.Unlock()
	return c.JSON(fiber.Map{"number": number})
}
