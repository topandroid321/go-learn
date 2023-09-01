package Tai

import (
	"github.com/gofiber/fiber/v2"
)

func GetTai(c *fiber.Ctx) error {
	return c.Status(200).SendString("Tai berhasil di get 💩!")
}
