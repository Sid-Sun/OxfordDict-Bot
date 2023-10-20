package query

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract"

	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/service"
	"go.uber.org/zap"
)

// Handler handles new queries
func Handler(bot *botAPI.BotAPI, update botAPI.Update, logger *zap.Logger, svc service.Service, adminChatID int64) {
	logger.Info("[Query] [Attempt]")

	// Treat first word in message as query and convert to lowercase
	query := strings.Fields(update.Message.Text)[0]

	if update.Message.IsCommand() {
		query_fields := strings.Fields(update.Message.CommandArguments())
		if len(query_fields) == 0 {
			reply := botAPI.NewMessage(update.Message.Chat.ID, "Invlaid usage, please use as: \"/en <word>\" ex: /en hello.")
			reply.ReplyToMessageID = update.Message.MessageID
			_, _ = bot.Send(reply)
			return
		}
		query = query_fields[0]
	}

	if !isLetters(query) {
		reply := botAPI.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Sorry, %s is not an English word.", query))
		reply.ReplyToMessageID = update.Message.MessageID
		_, _ = bot.Send(reply)
		return
	}

	definition, err := svc.GetDefinition(strings.ToLower(query))
	if err != nil {
		log := fmt.Sprintf("[%s] [Handler] [GetDefinition] %v", handler, err)
		logger.Error(log)
		if errors.Is(err, service.ErrForbidden) {
			reply := botAPI.NewMessage(update.Message.Chat.ID, "Sorry, quota reached, please try again later.")
			reply.ReplyToMessageID = update.Message.MessageID
			_, err = bot.Send(reply)
			return
		}
		adminMessage := botAPI.NewMessage(adminChatID, log)
		var reply botAPI.MessageConfig
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
			reply = botAPI.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Sorry, An internal error occurred. Please contact %s.", fmt.Sprintf("[%s](tg://user?id=%d)", name, adminChatID)))
			reply.ParseMode = "markdown"
		} else {
			reply = botAPI.NewMessage(update.Message.Chat.ID, "Sorry, An internal error occurred. Please try again later. Admins have been informed.")
		}
		reply.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(reply); err != nil {
			logger.Error(fmt.Sprintf("[%s] [Handler] [GetDefinition] [Error] [Send]", handler))
		}
		return
	}

	if definition.IsEmpty() {
		logger.Info(fmt.Sprintf("[%s] [Handler] [GetDefinition] [IsEmpty]", handler))
		reply := botAPI.NewMessage(update.Message.Chat.ID, "Sorry, I could not find definition for "+query)
		reply.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(reply); err != nil {
			logger.Error(fmt.Sprintf("[%s] [Handler] [GetDefinition] [IsEmpty] [Send]", handler))
		}
		logger.Info("[Query] [Success]")
		return
	}

	resp := contract.Response{
		APIResponse: definition,
		Query:       strings.ToLower(query),
	}

	formattedMessage := resp.GetFormatted(initialIndex)

	reply := botAPI.NewMessage(update.Message.Chat.ID, formattedMessage)
	reply.ReplyToMessageID = update.Message.MessageID

	numberOfDefinitions := definition.NumberOfDefinitions()

	if numberOfDefinitions > 1 {
		keyboardConfig := contract.KeyboardConfig{
			Current: initialIndex,
			Prev:    numberOfDefinitions - 1,
			Total:   numberOfDefinitions,
			Next:    1,
		}
		reply.ReplyMarkup = keyboardConfig.Keyboard()
	}

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

	logger.Info("[Query] [Success]")
}

func isLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) || unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
