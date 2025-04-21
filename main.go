package main

import (
	"log"
	"os"

	"github.com/Mario-Valente/kiwify-webhoock/cmd"
	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
	"github.com/Mario-Valente/kiwify-webhoock/internal/health"
	"github.com/Mario-Valente/kiwify-webhoock/internal/middleware"
	webhook "github.com/Mario-Valente/kiwify-webhoock/internal/webhoock/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config := config.NewConfig()
	app := fiber.New()

	app.Use(middleware.AuthMiddelware)

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

	go func() {
		if err := cmd.BotHandler(); err != nil {
			log.Println("Error starting bot:", err)
		}

	}()

	err := app.Listen(config.Port)
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}

}
