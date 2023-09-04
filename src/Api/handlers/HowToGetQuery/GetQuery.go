package HowToGetQuery

import (
	"context"
	GraphqlClient "go-learning/src/Utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// EXAMPLE HOW TO GET QUERY USING AUTH ADMIN
func ExampleGetUsingAdmin(c *fiber.Ctx) error {

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
func ExampleGetUsingUser(c *fiber.Ctx) error {

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

func ExampleGetPagination(c *fiber.Ctx) error {

	var query struct {
		Users []struct {
			ID any `json:"id"`
		} `graphql:"users(where: {id: {_eq: \"6cf18a5f-3333-428f-8d76-8abbdd4b168f\"}}, limit: 2, offset: 0, order_by: {id: desc})"`
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
