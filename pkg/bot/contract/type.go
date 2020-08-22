package contract

import (
	"fmt"
	"strconv"

	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"
)

// Response contains the the raw data for response
type Response struct {
	APIResponse api.Response
	Query       string
}

// GetFormatted returns formatted string for response
func (d Response) GetFormatted(index int) string {
	// Get required data
	definition := d.APIResponse.Results[0].LexicalEntries[0].Entries[0].Senses[index].Definitions[0]
	var example string
	if len(d.APIResponse.Results[0].LexicalEntries[0].Entries[0].Senses[index].Examples) > 0 {
		example = d.APIResponse.Results[0].LexicalEntries[0].Entries[0].Senses[index].Examples[0].Text
	}
	provider := d.APIResponse.Metadata.Provider
	language := d.APIResponse.Results[0].LexicalEntries[0].Language
	lexicalCategory := d.APIResponse.Results[0].LexicalEntries[0].LexicalCategory.Text

	// Format
	message := d.Query + " by " + provider
	if language != "" {
		message += "\n\nLanguage: \n" + language
	}
	message += "\n\nDefinition: \n" + definition
	message += "\n\nLexical Category: \n" + lexicalCategory
	if example != "" {
		message += "\n\nExamples: \n" + example
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
			botAPI.NewInlineKeyboardButtonData(fmt.Sprintf("%d/%d", k.Current, k.Total), "nah"),
			botAPI.NewInlineKeyboardButtonData("➡️", strconv.Itoa(k.Next)),
		),
	)
}
