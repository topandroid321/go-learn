package ApiRoutes

import (
	Billing "go-learning/src/Api/handlers/Billings"
	GetQuery "go-learning/src/Api/handlers/HowToGetQuery"
	"go-learning/src/Api/handlers/Index"
	"go-learning/src/Api/handlers/Users"
	Middleware "go-learning/src/Api/middleware"

	// "go-learning/src/Client"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRoutes(http *fiber.App) {
	http.Use(func(c *fiber.Ctx) error { return Middleware.CheckAuth(c) }) // auth middleware check
	http.Use(logger.New())                                                // logging

	http.Get("/", func(c *fiber.Ctx) error { return Index.Index(c) })
	http.Get("/users", func(c *fiber.Ctx) error { return Users.AddUsers(c) })
	http.Get("/get-admin", func(c *fiber.Ctx) error { return GetQuery.ExampleGetUsingAdmin(c) })
	http.Get("/get-user", func(c *fiber.Ctx) error { return GetQuery.ExampleGetUsingUser(c) })

	// Billings Routes
	Billings := http.Group("billing", func(c *fiber.Ctx) error { return c.Next() })
	Billings.Post("/create-customer", func(c *fiber.Ctx) error { return Billing.AddCustomer(c) })
	Billings.Get("/get-customer", func(c *fiber.Ctx) error { return Billing.GetCustomer(c) })
}
