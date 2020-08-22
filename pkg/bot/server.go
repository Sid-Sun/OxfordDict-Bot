package bot

import (
	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/router"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/service"
	"go.uber.org/zap"
)

// StartBot starts the bot, inits all the requited submodules and routine for shutdown
func StartBot(cfg config.Config, logger *zap.Logger) {
	svc := service.NewService(logger, cfg.API)
	ch := router.New(cfg.Bot, logger, svc).NewUpdateChan()

	logger.Info("[StartBot] Starting")
	ch.ListenAndServe()
}
