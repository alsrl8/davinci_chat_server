package middlewares

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"github.com/gofiber/fiber/v2"
)

func RedirectHTTPtoHTTPS(c *fiber.Ctx) error {
	runEnv := config.GetRunEnv()
	if runEnv == consts.Production {
		if c.Protocol() == "http" {
			return c.Redirect("https://"+c.Hostname()+c.OriginalURL(), 301)
		}
	}
	return c.Next()
}
