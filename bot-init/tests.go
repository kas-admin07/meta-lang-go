package botinit

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestInitializeBot(t *testing.T) {
	bot, err := InitializeBot("test-token")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if bot == nil {
		t.Errorf("Expected bot instance, got nil")
	}
}
