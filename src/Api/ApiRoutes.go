package ApiRoutes

import (
	Billing "go-learning/src/Api/handlers/Billings"
	"go-learning/src/Api/handlers/Index"
	"go-learning/src/Api/handlers/Tai"
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
	http.Get("/tai", func(c *fiber.Ctx) error { return Tai.GetTai(c) })

	Billings := http.Group("/billing", func(c *fiber.Ctx) error { return c.Next() })
	Billings.Post("/create-customer", func(c *fiber.Ctx) error { return Billing.AddCustomer(c) })
}
