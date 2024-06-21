package handlers

import (
	"context"
	"davinci-chat/user"
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

	service := user.NewUserService(context.Background())
	err := service.AddUser(
		context.Background(),
		&user.User{
			Name:     newUser.UserName,
			Password: newUser.Password,
		},
		newUser.UserEmail,
	)
	if err != nil {
		return err
	}

	return nil
}
