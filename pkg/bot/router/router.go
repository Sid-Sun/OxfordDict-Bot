package router

import (
	"fmt"
	"os"
	"strconv"

	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/handlers/callback"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/handlers/hello"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/handlers/query"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/service"
	"go.uber.org/zap"
)

// Updates contains update channel and defines method to listen and respond to them
type Updates struct {
	ch  botAPI.UpdatesChannel
	bot Bot
}

// ListenAndServe listens on the update channel and handles routing the update to handlers
func (u Updates) ListenAndServe() {
	u.bot.logger.Info(fmt.Sprintf("[Router] [ListenAndServe] Hi, I am %s", u.bot.bot.Self.String()))
	for update := range u.ch {
		update := update
		go func() {
			if update.CallbackQuery != nil {
				callback.Handler(u.bot.bot, update, u.bot.logger, u.bot.svc, u.bot.adminChatID)
				return
			}
			if update.Message == nil || update.Message.Text == "" {
				return
			}
			if cmd := update.Message.Command(); cmd != "" {
				switch cmd {
				case "start", "hello":
					hello.Handler(u.bot.bot, update, u.bot.logger)
					return
				}
			}
			query.Handler(u.bot.bot, update, u.bot.logger, u.bot.svc, u.bot.adminChatID)
		}()
	}
}

// Bot contains instances for functioning of the bot
type Bot struct {
	bot         *botAPI.BotAPI
	logger      *zap.Logger
	svc         service.Service
	adminChatID int64
}

// NewUpdateChan creates a new channel to get update
func (b Bot) NewUpdateChan() Updates {
	u := botAPI.NewUpdate(0)
	u.Timeout = 60
	ch, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	return Updates{ch: ch, bot: b}
}

// New returns a new instance of the router
func New(cfg config.BotConfig, logger *zap.Logger, svc service.Service) Bot {
	b, err := botAPI.NewBotAPI(cfg.Token())
	if err != nil {
		panic(err)
	}
	adminChatID, err := strconv.ParseInt(os.Getenv("ADMIN_CHAT_ID"), 10, 64)
	if err != nil {
		panic(err)
	}
	return Bot{
		bot:         b,
		logger:      logger,
		svc:         svc,
		adminChatID: adminChatID,
	}
}
