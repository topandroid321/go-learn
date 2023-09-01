package ApiRoutes

import (
	Billing "go-learning/src/Api/handlers/Billings"
	GetQuery "go-learning/src/Api/handlers/HowToGetQuery"
	"go-learning/src/Api/handlers/Index"
	"go-learning/src/Api/handlers/Users"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(http *fiber.App) {
	http.Get("/", func(c *fiber.Ctx) error { return Index.Index(c) })
	http.Get("/users", func(c *fiber.Ctx) error { return Users.AddUsers(c) })
	http.Get("/get-admin", func(c *fiber.Ctx) error { return GetQuery.ExampleGetUsingAdmin(c) })
	http.Get("/get-user", func(c *fiber.Ctx) error { return GetQuery.ExampleGetUsingUser(c) })

	// Billings Routes
	Billings := http.Group("billing", func(c *fiber.Ctx) error { return c.Next() })
	Billings.Post("/create-customer", func(c *fiber.Ctx) error { return Billing.AddCustomer(c) })
	Billings.Get("/get-customer", func(c *fiber.Ctx) error { return Billing.GetCustomer(c) })
}
