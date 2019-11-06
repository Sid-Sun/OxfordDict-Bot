package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func newThreeButtonInlineKeyboard(count string, data []string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️", data[0]),
			tgbotapi.NewInlineKeyboardButtonData(count, "nah"),
			tgbotapi.NewInlineKeyboardButtonData("➡️", data[1]),
		),
	)
}

func getMessage(response string, index int, query string) (string, bool) {
	definition := gjson.Get(response, "results.0.lexicalEntries.0.entries.0.senses."+strconv.Itoa(index)+".definitions.0").String()
	if definition == "" {
		return "", false
	}
	examples := gjson.Get(response, "results.0.lexicalEntries.0.entries.0.senses."+strconv.Itoa(index)+".examples.0.text").String()
	provider := gjson.Get(response, "metadata.provider").String()
	language := strings.ToUpper(gjson.Get(response, "results.0.lexicalEntries.0.language").String())
	lexicalCategory := gjson.Get(response, "results.0.lexicalEntries.0.lexicalCategory.text").String()
	message := query + " by " + provider
	if language != "" {
		message += "\n\nLanguage: " + language
	}
	message += "\n\nDefinition: \n" + definition
	message += "\n\nLexical Category: " + lexicalCategory
	if examples != "" {
		message += "\n\nExamples: \n" + examples
	}
	return message, true
}

func getDefinition(word string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://od-api.oxforddictionaries.com:443/api/v2/entries/en/"+strings.ToLower(word), nil)
	req.Header.Add("app_id", os.Getenv("APP_ID"))
	req.Header.Add("app_key", os.Getenv("APP_KEY"))
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
