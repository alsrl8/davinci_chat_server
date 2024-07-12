package handlers

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"davinci-chat/database"
	"davinci-chat/err_types"
	"davinci-chat/types"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var request types.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	db := database.GetDatabase()
	name, err := db.Login(request)
	if err != nil {
		switch {
		case errors.Is(err, err_types.ErrEmailPasswordNotMatch):
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email Password Not Match"})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to login"})
		}
	}

	email := request.UserEmail
	return c.JSON(fiber.Map{"message": "Login successful", "email": email, "name": name})
}

func getCookieDomain() string {
	runEnv := config.GetRunEnv()
	domain := ""
	switch runEnv {
	case consts.Development:
		domain = "localhost"
	case consts.Production:
		domain = "chat-dot-davinci-song.du.r.appspot.com"
	}
	return domain
}
