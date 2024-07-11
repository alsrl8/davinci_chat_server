package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func AddNewUser(c *fiber.Ctx) error {
	type NewUserRequest struct {
		UserName  string `json:"userName"`
		UserEmail string `json:"userEmail"`
		Password  string `json:"password"`
	}

	var newUser NewUserRequest
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	return nil
}
