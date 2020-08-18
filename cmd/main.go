package main

import (
	"github.com/sid-sun/sample-bot/cmd/config"
	"github.com/sid-sun/sample-bot/pkg/bot"
)

func main() {
	cfg := config.Load()
	initLogger(cfg.GetEnv())
	bot.StartBot(cfg, logger)
}
