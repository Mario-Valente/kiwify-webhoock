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

	result, err := Get(c.UserContext(), orderStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func Register(app *fiber.App) {
	app.Post("/webhook", post)
	app.Get("/webhook/:orderStatus", get)
}
