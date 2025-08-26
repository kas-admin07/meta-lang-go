package msghandler

// [go-msg-handler]
import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
)

func AddMessageHandler(bot *tgbotapi.BotAPI, action func(string) string) {
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	go func() {
		for update := range bot.GetUpdatesChan(tgbotapi.NewUpdate(0)) {
			if update.Message == nil || update.Message.Text == "" || update.Message.Text[0] == '/' {
				continue
			}
			url := urlRegex.FindString(update.Message.Text)
			if url != "" && IsValidURL(url) {
				result := action(url)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				bot.Send(msg)
			}
		}
	}()
}
