package Users

import (
	"github.com/gofiber/fiber/v2"
)

func AddUsers(c *fiber.Ctx) error {
	c.SendStatus(200)
	return c.SendString("Add Users !")
}
