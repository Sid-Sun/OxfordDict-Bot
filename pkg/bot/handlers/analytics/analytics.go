package analytics

import (
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/service"
)

func HandleAnalytics(chatID int64, svc service.Service) {
	svc.CollectAnalytics(chatID)
}
