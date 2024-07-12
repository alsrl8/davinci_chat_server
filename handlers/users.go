package handlers

import (
	"davinci-chat/database"
	"davinci-chat/err_types"
	"davinci-chat/types"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

func ValidateUser(c *fiber.Ctx) error {
	var validateUser types.ValidateUserRequest
	if err := c.BodyParser(&validateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}
	if err := validateUser.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db := database.GetDatabase()
	err := db.ValidateUser(validateUser)
	if err != nil {
		switch {
		case errors.Is(err, err_types.ErrUserExists):
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": validateUser})
}

func AddNewUser(c *fiber.Ctx) error {
	var newUser types.NewUserRequest
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	db := database.GetDatabase()
	err := db.AddUser(newUser)
	if err != nil {
		log.Fatal(err)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": newUser.UserName})
}
