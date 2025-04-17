package cmd

import (
	"context"
	"log"

	"github.com/Mario-Valente/kiwify-webhoock/internal/config"
	webhook "github.com/Mario-Valente/kiwify-webhoock/internal/webhoock/controllers"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func BotHandler() error {
	config := config.NewConfig()

	b, err := bot.New(config.TokenTelegram)
	if err != nil {
		return err
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Olá! Eu sou um bot feito para registar todas as movimentações da kiwify 🧠",
		})
		if err != nil {
			log.Println("Error sending message:", err)

		}
	})

	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "comandos disponiveis 😈 --> : \n😓 para pegar abandonos de carrinho: /abandoned \n💸para pegar todas vendas por categoria: /refused",
		})
		if err != nil {
			log.Println("Error sending message:", err)
		}
	})

	b.RegisterHandler(bot.HandlerTypeMessageText, "/abandoned", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		chatID := update.Message.Chat.ID

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "🔍 Buscando abandonos de carrinho, aguarde um momento...",
		})
		if err != nil {
			log.Println("Erro ao enviar mensagem inicial:", err)
			return
		}

		purchases, err := webhook.GetAllAbandoned(ctx)
		if err != nil {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❌ Erro ao buscar dados: " + err.Error(),
			})
			return
		}

		message := "🛒 Abandonos de carrinho encontrados:\n\n"
		for i, p := range purchases {
			message += formatAbandoned(p, i+1)
		}

		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   message,
		})
		if err != nil {
			log.Println("Erro ao enviar dados:", err)
		}
	})

	b.RegisterHandler(bot.HandlerTypeMessageText, "/refused", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		chatID := update.Message.Chat.ID
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "🔍 Buscando vendas recusadas, aguarde um momento...",
		})
		if err != nil {
			log.Println("Erro ao enviar mensagem inicial:", err)
			return
		}
		purchases, err := webhook.GetAllByStatus(ctx, "refused")
		if err != nil {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❌ Erro ao buscar dados: " + err.Error(),
			})
			return
		}
		message := "🛒 Vendas recusadas encontradas:\n\n"
		for i, p := range purchases {
			message += formatPurchase(p, i+1)
		}
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   message,
		})
		if err != nil {
			log.Println("Erro ao enviar dados:", err)
		}
	})

	b.Start(context.Background())
	return nil

}
