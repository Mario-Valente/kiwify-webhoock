package main

import (
	"github.com/Mario-Valente/kiwify-webhoock/internal/health"
	webhook "github.com/Mario-Valente/kiwify-webhoock/internal/webhoock/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	health.Register(app)
	webhook.Register(app)

	app.Listen(":3000")

}
