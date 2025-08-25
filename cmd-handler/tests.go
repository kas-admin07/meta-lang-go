package cmdhandler

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestAddCommand(t *testing.T) {
	bot, _ := tgbotapi.NewBotAPI("test-token")
	updates := make(chan tgbotapi.Update, 1)
	bot.GetUpdatesChan = func(_ tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
		return updates
	}
	bot.Send = func(c tgbotapi.Chattable) (tgbotapi.Message, error) {
		msg := c.(tgbotapi.MessageConfig)
		if msg.Text != "Welcome!" {
			t.Errorf("Expected response 'Welcome!', got '%s'", msg.Text)
		}
		return tgbotapi.Message{}, nil
	}

	AddCommand(bot, "start", "Welcome!")
	updates <- tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{ID: 1},
			Text: "/start",
		},
	}
}
