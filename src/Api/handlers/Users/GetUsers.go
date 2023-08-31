package Users

import (
	"context"
	GraphqlClient "go-learning/src/Utils"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	var query struct {
		Me struct {
			Name string
		}
	}

	client := GraphqlClient.Client
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
		panic(err)
	}
	return c.SendString("users")
}
