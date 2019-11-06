package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tidwall/gjson"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		fmt.Println(err.Error())
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err.Error())
	}
	threads, err := strconv.Atoi(os.Getenv("NUMBER_OF_THREADS"))
	if err != nil {
		fmt.Println(err.Error())
		threads = 2
	}
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go getUpdates(bot, updates)
	}
}

func getUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		go handleUpdate(bot, update)
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update)  {
	if update.CallbackQuery != nil {
		response, err := getDefinition(strings.ToLower(strings.Fields(update.CallbackQuery.Message.ReplyToMessage.Text)[0]))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		sensesCount, err := strconv.Atoi(gjson.Get(response, "results.0.lexicalEntries.0.entries.0.senses.#").String())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data, err := strconv.Atoi(update.CallbackQuery.Data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		current := data
		next := data + 1
		previous := data - 1
		if sensesCount-1 == data {
			next = 0
		}
		if 0 == data {
			next = data + 1
			previous = sensesCount - 1
		}
		message, _ := getMessage(response, current, strings.Fields(update.CallbackQuery.Message.ReplyToMessage.Text)[0])
		temp := tgbotapi.NewEditMessageText(int64(update.CallbackQuery.From.ID), update.CallbackQuery.Message.MessageID, message)
		newInlineKeyboardMarkup := newThreeButtonInlineKeyboard(strconv.Itoa(current+1)+"/"+strconv.Itoa(sensesCount), []string{strconv.Itoa(previous), strconv.Itoa(next)})
		temp.ReplyMarkup = &newInlineKeyboardMarkup
		mess, err := bot.Send(temp)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(mess)
		return
	}
	if update.Message == nil { // ignore any non-Message Updates
		return
	}
	response, err := getDefinition(strings.ToLower(strings.Fields(update.Message.Text)[0]))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var message string
	var successful bool
	if message, successful = getMessage(response, 0, strings.Fields(update.Message.Text)[0]); !successful {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, I could not find definition for "+update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	sensesCount := gjson.Get(response, "results.0.lexicalEntries.0.entries.0.senses.#").String()
	numberOfSenses, _ := strconv.Atoi(sensesCount)
	numberOfSenses--
	msg.ReplyMarkup = newThreeButtonInlineKeyboard("1/"+sensesCount, []string{strconv.Itoa(numberOfSenses), "1"})
	if _, err := bot.Send(msg); err != nil {
		fmt.Println(err.Error())
	}
}