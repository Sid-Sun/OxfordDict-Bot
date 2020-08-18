package bot

import (
	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/router"
	"go.uber.org/zap"
)

// StartBot starts the bot, inits all the requited submodules and routine for shutdown
func StartBot(cfg config.Config, logger *zap.Logger) {
	ch := router.New(cfg.Bot, logger).NewUpdateChan()

	logger.Info("[StartBot] Started Bot")
	ch.ListenAndServe()
}
