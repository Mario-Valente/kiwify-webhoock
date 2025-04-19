package cmd

import (
	"context"
	"fmt"
	"log"
	"regexp"

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

	re := regexp.MustCompile(`^(\/refused|\/waiting_payment|\/refunded|\/chargedback)`)

	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, re, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		chatID := update.Message.Chat.ID
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "🔍 Buscando vendas por status, aguarde um momento...",
		})
		if err != nil {
			log.Println("Erro ao enviar mensagem inicial:", err)
			return
		}

		command := update.Message.Text
		var status string
		switch command {
		case "/refused":
			status = "refused"
		case "/waiting_payment":
			status = "waiting_payment"
		case "/refunded":
			status = "refunded"
		case "/chargedback":
			status = "chargedback"
		default:
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❌ Comando inválido.",
			})
			return
		}
		purchases, err := webhook.GetAllByStatus(ctx, status)
		if err != nil {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❌ Erro ao buscar dados: " + err.Error(),
			})
			return
		}
		message := fmt.Sprintf("🛒 Vendas com status %s encontradas:\n\n", status)
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

	rePaymentMethod := regexp.MustCompile(`^(\/credit_card|\/pix|\/boleto)`)

	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, rePaymentMethod, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		chatID := update.Message.Chat.ID
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "🔍 Buscando vendas por metodo de pagamento, aguarde um momento...",
		})
		if err != nil {
			log.Println("Erro ao enviar mensagem inicial:", err)
			return
		}

		command := update.Message.Text
		var payment_method string
		switch command {
		case "/credit_card":
			payment_method = "credit_card"
		case "/pix":
			payment_method = "pix"
		case "/boleto":
			payment_method = "boleto"
		default:
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❌ Comando inválido.",
			})
			return
		}
		purchases, err := webhook.GetAllByPaymentMethod(ctx, payment_method)
		if err != nil {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❌ Erro ao buscar dados: " + err.Error(),
			})
			return
		}
		message := fmt.Sprintf("🛒 Vendas com payment_method %s encontradas:\n\n", payment_method)
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
