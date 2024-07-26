package handlers

import (
	"davinci-chat/database"
	"davinci-chat/err_types"
	"davinci-chat/types"
	"davinci-chat/utils"
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

	token, err := utils.MakeJwt(name, request.UserEmail, false)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	cookie := utils.SetJwtCookie(token)
	c.Cookie(cookie)

	return c.JSON(fiber.Map{"message": "Login successful", "email": request.UserEmail, "name": name, "isGuest": false})
}

func AutoLoginHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		err := setGuestToken(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to generate guest"})
		}
		return c.JSON(fiber.Map{"name": "Guest", "email": "", "isGuest": true})
	}

	name, email, isGuest, err := utils.DecodeJwt(token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode token"})
	}

	return c.JSON(fiber.Map{"name": name, "email": email, "isGuest": isGuest})
}

func setGuestToken(c *fiber.Ctx) error {
	_token, err := utils.MakeJwt("Guest", "", true)
	if err != nil {
		return err
	}
	cookie := utils.SetJwtCookie(_token)
	c.Cookie(cookie)
	return nil
}
