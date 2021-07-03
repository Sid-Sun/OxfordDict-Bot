package bot

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/router"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/service"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/store"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

// StartBot starts the bot, inits all the requited submodules and routine for shutdown
func StartBot(cfg config.Config, logger *zap.Logger) {
	rdc, err := store.NewRedisClientConfig(logger, cfg.Redis).GetClient()
	if err != nil {
		panic(err)
	}
	str := store.NewStore(store.NewRedisService(rdc, logger))
	svc := service.NewService(logger, &cfg.API, str)
	ch := router.New(cfg.Bot, logger, svc).NewUpdateChan()

	logger.Info("[StartBot] Starting")
	go ch.ListenAndServe()

	gracefulShutdown(rdc, logger)
}

func gracefulShutdown(rdc *redis.Client, logger *zap.Logger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	logger.Info("Attempting GracefulShutdown")
	go func() {
		if err := rdc.Close(); err != nil {
			logger.Error(fmt.Sprintf("[GracefulShutdown] [Shutdown] [Redis]: %s", err.Error()))
			panic(err)
		}
	}()
}
