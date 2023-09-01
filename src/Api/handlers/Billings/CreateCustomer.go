package Billings

import (
	"go-learning/src/Interfaces"
	stripeClient "go-learning/src/Utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func AddCustomer(c *fiber.Ctx) error {

	newCustommer := Interfaces.AddCustomer{}

	if err := c.BodyParser(&newCustommer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validasi data input
	if newCustommer.Name == "" || newCustommer.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nama, Email, Required",
		})
	}

	stripeClient.InitStripe()
	params := &stripe.CustomerParams{
		Email:       stripe.String(newCustommer.Email),
		Name:        stripe.String(newCustommer.Name),
		Description: stripe.String(newCustommer.Description),
	}
	cus, _ := customer.New(params)

	return c.JSON(cus)
}
