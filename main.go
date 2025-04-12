package main

import (
	"log"
	"os"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
	"github.com/Mario-Valente/kiwify-webhoock/internal/health"
	webhook "github.com/Mario-Valente/kiwify-webhoock/internal/webhoock/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config := config.NewConfig()
	app := fiber.New()

	health.Register(app)
	webhook.Register(app)

	app.Use(logger.New(
		logger.Config{
			Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error} | " + config.ServiceName + "\n",
			TimeFormat: "15:04:05",
			TimeZone:   "Local",
			Output:     os.Stdout,
		},
	))

	err := app.Listen(config.Port)
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}

}
