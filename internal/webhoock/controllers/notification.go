package webhook

import (
	"context"
	"log"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
	"github.com/go-telegram/bot"
)

func SendTelegramMessage(ctx context.Context, message string) error {

	config := config.NewConfig()

	b, err := bot.New(config.TokenTelegram)
	if err != nil {
		log.Println("Error creating bot:", err)
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: config.ChatID,
		Text:   message,
	})
	if err != nil {
		log.Println("Error sending message:", err)
	}

	return nil
}
