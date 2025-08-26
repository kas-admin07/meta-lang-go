package msghandler

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestAddMessageHandler(t *testing.T) {
	bot, _ := tgbotapi.NewBotAPI("test-token")
	updates := make(chan tgbotapi.Update, 1)
	bot.GetUpdatesChan = func(_ tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
		return updates
	}
	bot.Send = func(c tgbotapi.Chattable) (tgbotapi.Message, error) {
		msg := c.(tgbotapi.MessageConfig)
		if msg.Text != "Shortened URL" {
			t.Errorf("Expected response 'Shortened URL', got '%s'", msg.Text)
		}
		return tgbotapi.Message{}, nil
	}

	AddMessageHandler(bot, func(url string) string {
		if url != "https://example.com" {
			t.Errorf("Expected URL 'https://example.com', got '%s'", url)
		}
		return "Shortened URL"
	})
	updates <- tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{ID: 1},
			Text: "Check this: https://example.com",
		},
	}
}
