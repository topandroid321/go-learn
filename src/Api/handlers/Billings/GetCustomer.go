package Billings

import (
	stripeClient "go-learning/src/Utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/customer"
)

type GetCustomerId struct {
	Customer_id string `json:"customer_id"`
}

func GetCustomer(c *fiber.Ctx) error {
	stripeClient.InitStripe()

	custommer := GetCustomerId{}
	if err := c.BodyParser(&custommer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "Cannot parse JSON",
		})
	}

	if custommer.Customer_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "customer_id, is Required",
		})
	}

	cus, _ := customer.Get(custommer.Customer_id, nil)

	response := fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":       true,
			"messages":    "success-get-data-customer",
			"customer_id": cus.ID,
			"email":       cus.Email,
			"name":        cus.Name,
			"description": cus.Description,
		},
	}

	return c.JSON(response)
}
