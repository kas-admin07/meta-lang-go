package cmdhandler

// [go-cmd-handler]
import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddCommand(bot *tgbotapi.BotAPI, command, response string) {
	go func() {
		for update := range bot.GetUpdatesChan(tgbotapi.NewUpdate(0)) {
			if update.Message == nil || update.Message.Command() != command {
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(msg)
		}
	}()
}
