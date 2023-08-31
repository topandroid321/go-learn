package ApiRoutes

import (
	"go-learning/src/Api/handlers/Index"
	"go-learning/src/Api/handlers/Users"
	"go-learning/src/Client"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes() {
	http := Client.App

	http.Get("/", func(c *fiber.Ctx) error { return Index.Index(c) })
	http.Get("/users", func(c *fiber.Ctx) error { return Users.AddUsers(c) })
}
