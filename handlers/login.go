package handlers

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	//type LoginRequest struct {
	//	RefreshToken string `json:"refreshToken"`
	//}
	//
	//var request LoginRequest
	//if err := c.BodyParser(&request); err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	//}
	//
	//idToken := c.Get("Authorization")
	//if idToken == "" {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid JWT"})
	//}
	//
	//idToken = strings.TrimPrefix(idToken, "Bearer ")
	//
	//decodedToken, err := auth.FirebaseAuth.VerifyIDToken(context.Background(), idToken)
	//if err != nil {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid JWT"})
	//}
	//
	//email := decodedToken.Claims["email"].(string)
	//userService := user.NewUserService(context.Background())
	//getUser, err := userService.GetUser(context.Background(), email)
	//if err != nil {
	//	return err
	//}
	//
	//userName := getUser.Name
	//
	//c.Cookie(&fiber.Cookie{
	//	Name:     "idToken",
	//	Value:    idToken,
	//	Expires:  time.Now().Add(time.Hour),
	//	HTTPOnly: true,
	//	Secure:   true,
	//	SameSite: fiber.CookieSameSiteNoneMode,
	//	Path:     "/",
	//	Domain:   getCookieDomain(),
	//})
	//
	//c.Set("X-Id-Token", idToken)
	//return c.JSON(fiber.Map{"message": "Login successful", "email": email, "name": userName})
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
