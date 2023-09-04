package HowToGetQuery

import (
	"context"
	"encoding/json"
	"fmt"
	"go-learning/src/Utils/Functions"
	"go-learning/src/Utils/GraphqlClient"

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

// EXAMPLE HOW TO GET QUERY USING FILTER PAGINATION (LIMIT, OFFSET)
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

// EXAMPLE HOW TO GET QUERY USING WHERE FILTER
func ExampleGetWhere(c *fiber.Ctx) error {

	email := "%" + c.Query("email") + "%"
	// limit := c.Query("limit")
	limit := Functions.IsEmpty(c.Query("limit"), 2)
	offset := Functions.IsEmpty(c.Query("offset"), 0)

	query := fmt.Sprintf(`
		query {
			users(where: { email: { _ilike: "%s" }}, limit: %d, offset: %d) {
				id
				email
			}
		}
	`, email, limit, offset)

	log.Debug(query)

	// Create a struct to unmarshal the response data into
	var res struct {
		Users []struct {
			ID    string `json:"id"`
			Email string `json:"email"`
		} `json:"users"`
	}

	// Create a GraphQL client
	client := GraphqlClient.CreateAdmin()

	// Execute the GraphQL query and capture the raw response
	raw, err := client.ExecRaw(context.Background(), query, map[string]interface{}{})
	if err != nil {
		panic(err)
	}

	// Unmarshal the raw response into the 'res' struct
	if err := json.Unmarshal(raw, &res); err != nil {
		panic(err)
	}

	// Return the 'res' struct as JSON response
	return c.JSON(res)
}
