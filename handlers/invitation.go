package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetUserEmailsByNameHandler(c *fiber.Ctx) error {
	userName := c.Query("name")

	var emails []string
	for conn, valid := range connections {
		if !valid {
			continue
		} else if connUserMap[conn].UserName != userName {
			continue
		}

		emails = append(emails, connUserMap[conn].UserEmail)
	}

	if len(emails) == 0 {
		return c.JSON(fiber.Map{"emails": []string{}})
	}

	return c.JSON(fiber.Map{"emails": emails})
}
