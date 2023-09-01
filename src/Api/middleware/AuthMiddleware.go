package Middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	if c.Path() == "/" {
		return c.Next()
	} else {
		header := c.GetReqHeaders()
		token := header["Authorization"]
		// log.Info("token : " + token)
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}
		return c.Next()
	}
}
