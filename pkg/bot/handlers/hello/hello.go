package hello

import (
	"fmt"

	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

// Handler handles start or hello requests
func Handler(bot *botAPI.BotAPI, update botAPI.Update, logger *zap.Logger) {
	logger.Info("[Hello] [Attempt]")

	reply := botAPI.NewMessage(update.Message.Chat.ID, greeting)

	_, err := bot.Send(reply)
	if err != nil {
		logger.Error(fmt.Sprintf("[%s] [Handler] [Send]", handler))
		return
	}

	logger.Info("[Hello] [Success]")
}
