package webhook

import (
	"github.com/Mario-Valente/kiwify-webhoock/internal/models"
	"github.com/gofiber/fiber/v2"
)

func post(c *fiber.Ctx) error {
	body := new(models.Purchase)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := Post(c.UserContext(), body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process webhook",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)

}

func get(c *fiber.Ctx) error {
	orderStatus := c.Params("orderStatus")
	if orderStatus == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer name is required",
		})
	}

	result, err := GetAllByStatus(c.UserContext(), orderStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func postAbandoned(c *fiber.Ctx) error {
	body := new(models.Abandoned)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := PostAbandoned(c.UserContext(), body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process webhook",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)

}

func getAllAbandoned(c *fiber.Ctx) error {

	result, err := GetAllAbandoned(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
func getAllByPaymentMethod(c *fiber.Ctx) error {
	paymentMethod := c.Params("paymentMethod")
	if paymentMethod == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Payment method is required",
		})
	}
	result, err := GetAllByPaymentMethod(c.UserContext(), paymentMethod)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func Register(app *fiber.App) {
	app.Post("/webhook", post)
	app.Post("/webhook/abandoned", postAbandoned)
	app.Get("/webhook/abandoned", getAllAbandoned)
	app.Get("/webhook/:orderStatus", get)
	app.Get("/webhook/payment/:paymentMethod", getAllByPaymentMethod)
}
