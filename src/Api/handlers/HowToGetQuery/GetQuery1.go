package HowToGetQuery

import (
	"context"
	"go-learning/src/Utils/GraphqlClient"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// EXAMPLE HOW TO GET QUERY USING AUTH ADMIN
func ExampleGetUsingAdmin1(c *fiber.Ctx) error {

	var query struct {
		Users []struct {
			ID any `json:"id"`
		} `json:"users"`
	}

	client := GraphqlClient.CreateAdmin()
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Error(err)
		log.Debug(query)
		return c.Status(fiber.StatusBadGateway).SendString("Something went wrong : " + err.Error())
	}

	log.Debug(query.Users)
	return c.JSON(query)
}

// EXAMPLE HOW TO GET QUERY USING AUTH USER
func ExampleGetUsingUser1(c *fiber.Ctx) error {

	var query struct {
		Users []struct {
			Username any `json:"username"`
		} `json:"users"`
	}

	headers := c.GetReqHeaders()
	token := headers["Authorization"]

	client := GraphqlClient.CreateClient(token)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Error(err)
		log.Debug(query)
		return c.Status(fiber.StatusBadGateway).SendString("Something went wrong : " + err.Error())
	}

	log.Debug(query.Users)
	return c.JSON(query)
}
