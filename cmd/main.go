package main

import (
	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot"
)

func main() {
	cfg := config.Load()
	initLogger(cfg.GetEnv())
	bot.StartBot(cfg, logger)
}
