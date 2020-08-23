package callback

import (
	"fmt"
	"strconv"
	"strings"

	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/service"
	"go.uber.org/zap"
)

// Handler handles callback queries
func Handler(bot *botAPI.BotAPI, update botAPI.Update, logger *zap.Logger, svc service.Service, adminChatID int64) {
	logger.Info("[Callback] [Attempt]")

	// Defer callback query answer
	defer func() {
		newCallBackConfig := botAPI.NewCallback(update.CallbackQuery.ID, "")
		_, err := bot.AnswerCallbackQuery(newCallBackConfig)
		if err != nil {
			log := fmt.Sprintf("[%s] [Handler] [AnswerCallbackQuery] %v", handler, err)
			logger.Error(log)
			adminMessage := botAPI.NewMessage(adminChatID, log)
			if _, err := bot.Send(adminMessage); err != nil {
				logger.Error(fmt.Sprintf("[%s] [Handler] [AnswerCallbackQuery] [Error] [Send] %v", handler, err))
			}
			return
		}
	}()

	query := strings.Fields(update.CallbackQuery.Message.ReplyToMessage.Text)[0]
	definition, err := svc.GetDefinition(strings.ToLower(query))
	if err != nil {
		log := fmt.Sprintf("[%s] [Handler] [GetDefinition] %v", handler, err)
		logger.Error(log)
		adminMessage := botAPI.NewMessage(adminChatID, log)
		var reply botAPI.EditMessageTextConfig
		if _, err := bot.Send(adminMessage); err != nil {
			logger.Error(fmt.Sprintf("[%s] [Handler] [GetDefinition] [Error] [Admin] [Send] %v", handler, err))
			c, err := bot.GetChat(botAPI.ChatConfig{
				ChatID: adminChatID,
			})
			if err != nil {
				logger.Error(fmt.Sprintf("[%s] [Handler] [GetDefinition] [Error] [Admin] [Send] [GetChat] %v", handler, err))
			}
			name := "Admin"
			if c.UserName != "" {
				name = c.UserName
			} else if c.FirstName != "" {
				name = c.FirstName
			}
			reply = botAPI.NewEditMessageText(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, fmt.Sprintf("Sorry, An internal error occurred. Please contact %s.", fmt.Sprintf("[%s](tg://user?id=%d)", name, adminChatID)))
			reply.ParseMode = "markdown"
		} else {
			reply = botAPI.NewEditMessageText(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, "Sorry, An internal error occurred. Please try again later. Admins have been informed.")
		}
		if _, err := bot.Send(reply); err != nil {
			logger.Error(fmt.Sprintf("[%s] [Handler] [GetDefinition] [Error] [Send]", handler))
		}
		return
	}

	if update.CallbackQuery.Data == "nah" {
		logger.Info("[Callback] [Success]")
		return
	}

	data, err := strconv.Atoi(update.CallbackQuery.Data)
	if err != nil {
		logger.Error(fmt.Sprintf("[%s] [Handler] [Atoi]", handler))
		return
	}

	numberOfDefinitions := definition.NumberOfDefinitions()

	current := data
	next := data + 1
	previous := data - 1
	if numberOfDefinitions-1 == data {
		next = 0
	}
	if data == 0 {
		previous = numberOfDefinitions - 1
	}

	keyboard := contract.KeyboardConfig{
		Total:   numberOfDefinitions,
		Current: current + 1,
		Next:    next,
		Prev:    previous,
	}.Keyboard()

	resp := contract.Response{
		APIResponse: definition,
		Query:       strings.ToLower(query),
	}

	formattedMessage := resp.GetFormatted(current)

	reply := botAPI.NewEditMessageText(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, formattedMessage)
	reply.ReplyMarkup = &keyboard

	_, err = bot.Send(reply)
	if err != nil {
		log := fmt.Sprintf("[%s] [Handler] [Send] %v", handler, err)
		logger.Error(log)
		adminMessage := botAPI.NewMessage(adminChatID, log)
		if _, err := bot.Send(adminMessage); err != nil {
			logger.Error(fmt.Sprintf("[%s] [Handler] [Send] [Error] [Send] %v", handler, err))
		}
		return
	}

	logger.Info("[Callback] [Success]")
}
