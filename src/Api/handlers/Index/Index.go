package Index

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	c.SendStatus(200)
	return c.SendString("Hello world !")
}
