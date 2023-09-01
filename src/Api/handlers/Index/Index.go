package Index

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Status(200).SendString("API Running Well!")
}
