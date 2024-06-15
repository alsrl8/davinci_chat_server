package handlers

import (
	"context"
	"davinci-chat/auth"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"time"
)

func LoginHandler(c *fiber.Ctx) error {
	type LoginRequest struct {
		RefreshToken string `json:"refreshToken"`
	}

	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	idToken := c.Get("Authorization")
	if idToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid JWT"})
	}

	idToken = strings.TrimPrefix(idToken, "Bearer ")

	decodedToken, err := auth.FirebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT"})
	}

	log.Printf("User authenticated with UID: %s", decodedToken.UID)

	c.Cookie(&fiber.Cookie{
		Name:     "idToken",
		Value:    idToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
		Domain:   "localhost",
	})

	return c.JSON(fiber.Map{"message": "Login successful"})
}
