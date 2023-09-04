package ApiRoutes

import (
	Billing "go-learning/src/Api/handlers/Billings"
	HowToGetQuery "go-learning/src/Api/handlers/HowToGetQuery"
	"go-learning/src/Api/handlers/Index"
	"go-learning/src/Api/handlers/Users"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(http *fiber.App) {
	http.Get("/", func(c *fiber.Ctx) error { return Index.Index(c) })
	http.Get("/users", func(c *fiber.Ctx) error { return Users.AddUsers(c) })
	http.Get("/get-admin", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetUsingAdmin(c) })
	http.Get("/get-user", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetUsingUser(c) })
	http.Get("/get-pagination", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetPagination(c) })
	http.Get("/get-where", func(c *fiber.Ctx) error { return HowToGetQuery.ExampleGetWhere(c) })

	// Billings Routes
	Billings := http.Group("billing", func(c *fiber.Ctx) error { return c.Next() })
	Billings.Post("/create-customer", func(c *fiber.Ctx) error { return Billing.AddCustomer(c) })
	Billings.Get("/get-customer", func(c *fiber.Ctx) error { return Billing.GetCustomer(c) })

	// Users Routes Mysql
	http.Get("/test", func(c *fiber.Ctx) error { return Users.TestConnection(c) })
	http.Get("/users/:id", func(c *fiber.Ctx) error { return Users.GetUserDetails(c) })
}
