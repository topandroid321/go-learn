package Users

import (
	"context"
	GraphqlClient "go-learning/src/Utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetUsers(c *fiber.Ctx) error {

	var query struct {
		users struct {
			id any
		}
	}

	headers := c.GetReqHeaders()
	token := headers["Authorization"]

	client := GraphqlClient.CreateClient(token)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Error(err)
		log.Debug(query)
		return c.Status(500).SendString("Something went wrong : " + err.Error())
	}
	log.Debug(query.users.id)
	return c.SendString("hello")
}
