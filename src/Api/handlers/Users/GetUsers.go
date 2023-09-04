package Users

import (
	"context"
	"fmt"
	"go-learning/src/Utils/GraphqlClient"

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
		return c.Status(fiber.StatusOK).SendString("Something went wrong : " + err.Error())
	}
	log.Debug(query.users.id)
	return c.SendString(fmt.Sprint(query.users.id))
}
