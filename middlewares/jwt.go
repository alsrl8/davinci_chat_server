package middlewares

import (
	"context"
	"davinci-chat/auth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func JWTMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("idToken")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid JWT"})
	}

	token = strings.TrimPrefix(token, "Bearer ")

	decodedToken, err := auth.FirebaseAuth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT"})
	}

	c.Locals("uid", decodedToken.UID)
	return c.Next()
}
