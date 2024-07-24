package handlers

import (
	"davinci-chat/utils"
	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(c *fiber.Ctx) error {
	cookie := utils.MakeJwtCookie("")
	c.Cookie(cookie)
	return c.JSON(fiber.Map{"message": "Logout successful"})
}
