package Billings

import (
	"go-learning/src/Interfaces"
	"go-learning/src/Utils/StripeClient"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func AddCustomer(c *fiber.Ctx) error {

	newCustommer := Interfaces.AddCustomer{}

	if err := c.BodyParser(&newCustommer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "Cannot parse JSON",
		})
	}

	if newCustommer.Name == "" || newCustommer.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"error":      "Nama, Email, is eequired",
		})
	}

	StripeClient.InitStripe()
	params := &stripe.CustomerParams{
		Email:       stripe.String(newCustommer.Email),
		Name:        stripe.String(newCustommer.Name),
		Description: stripe.String(newCustommer.Description),
	}
	cus, _ := customer.New(params)

	response := fiber.Map{
		"statusCode": 200,
		"data": fiber.Map{
			"valid":       true,
			"messages":    "success-create-customer",
			"customer_id": cus.ID,
		},
	}

	return c.JSON(response)
}
