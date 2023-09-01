package ApiRoutes

import (
	Billing "go-learning/src/Api/handlers/Billings"
	"go-learning/src/Api/handlers/Index"
	"go-learning/src/Api/handlers/Users"
	"go-learning/src/Client"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes() {
	http := Client.App

	http.Get("/", func(c *fiber.Ctx) error { return Index.Index(c) })
	http.Get("/users", func(c *fiber.Ctx) error { return Users.AddUsers(c) })

	Billings := http.Group("/billing", func(c *fiber.Ctx) error { return c.Next() })
	Billings.Post("/create-customer", func(c *fiber.Ctx) error { return Billing.AddCustomer(c) })
}
