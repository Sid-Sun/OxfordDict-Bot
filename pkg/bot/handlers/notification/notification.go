package notification

import (
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/service"
	"go.uber.org/zap"
)

func HandleAnalyticsAndNotify(bot *botAPI.BotAPI, chatID int64, logger *zap.Logger, svc service.Service, adminChatID int64) {
	if !svc.IsUserNotified(chatID) {
		msg := botAPI.NewMessage(chatID, "Hi, this bot has been deprecated in favour of @googlebot; please contact @SidSun if you have any concerns")
		bot.Send(msg)
	}
}

