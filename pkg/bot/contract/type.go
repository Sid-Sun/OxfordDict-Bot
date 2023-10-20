package contract

import (
	"fmt"
	"strconv"
	"unicode"

	botAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"
)

// Response contains the the raw data for response
type Response struct {
	APIResponse api.Response
	Query       string
}

func formatAsSentence(s string) string {
	r := []rune(s)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}

// GetFormatted returns formatted string for response
func (d Response) GetFormatted(index int) string {
	// Get required data
	var definition string
	if d.APIResponse.NumberOfDefinitions() > 0 {
		definition = formatAsSentence(d.APIResponse.Results[index].Definition)
	}
	var examples string
	if len(d.APIResponse.Results[index].Examples) > 0 {
		for i, example := range d.APIResponse.Results[index].Examples {
			examples += fmt.Sprintf("\n%d. %s", i+1, formatAsSentence(example))
		}
	}

	lexicalCategory := d.APIResponse.Results[index].PartOfSpeech

	// Format
	message := fmt.Sprintf("%s by WordsAPI", d.APIResponse.Word)
	if definition != "" {
		message += "\n\nDefinition: \n" + definition
	}
	message += "\n\nLexical Category: \n" + lexicalCategory
	if examples != "" {
		message += "\n\nExample(s): " + examples
	}

	return message
}

// KeyboardConfig defines config for inline keyboard
type KeyboardConfig struct {
	Total   int
	Current int
	Next    int
	Prev    int
}

// Keyboard returns a new inline keyboard with required data
func (k KeyboardConfig) Keyboard() botAPI.InlineKeyboardMarkup {
	return botAPI.NewInlineKeyboardMarkup(
		botAPI.NewInlineKeyboardRow(
			botAPI.NewInlineKeyboardButtonData("⬅️", strconv.Itoa(k.Prev)),
			botAPI.NewInlineKeyboardButtonData(fmt.Sprintf("%d/%d", k.Current+1, k.Total), "nah"),
			botAPI.NewInlineKeyboardButtonData("➡️", strconv.Itoa(k.Next)),
		),
	)
}
