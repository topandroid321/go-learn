package Tai

import (
	"context"
	GraphqlClient "go-learning/src/Utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetTai(c *fiber.Ctx) error {

	var query struct {
		Users []struct {
			ID any `json:"id"`
		} `json:"users"`
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

	log.Debug(query.Users)
	return c.JSON(query)
}
