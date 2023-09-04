package Billings

import (
	"go-learning/src/Utils/StripeClient"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/customer"
)

type GetCustomerId struct {
	Customer_id string `json:"customer_id"`
}

func GetCustomer(c *fiber.Ctx) error {
	StripeClient.InitStripe()

	custommer := GetCustomerId{}
	if err := c.BodyParser(&custommer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if custommer.Customer_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "customer_id, is Required",
		})
	}

	StripeClient.InitStripe()
	cus, _ := customer.Get(custommer.Customer_id, nil)

	response := fiber.Map{
		"statusCode": 200,
		"data": fiber.Map{
			"valid":       true,
			"messages":    "Success Get Data Customer",
			"customer_id": cus.ID,
			"email":       cus.Email,
			"name":        cus.Name,
			"description": cus.Description,
		},
	}

	return c.JSON(response)
}
