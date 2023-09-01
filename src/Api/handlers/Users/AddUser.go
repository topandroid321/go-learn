package Users

import (
	"github.com/gofiber/fiber/v2"
)

func AddUsers(c *fiber.Ctx) error {
	return c.Status(200).SendString("Add Users !")
}
