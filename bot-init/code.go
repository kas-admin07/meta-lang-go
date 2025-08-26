package botinit

// [go-bot-init]
import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitializeBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return bot, nil
}
