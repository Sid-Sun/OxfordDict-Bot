package config

// BotConfig contains the config for creating a new bot
type BotConfig struct {
	tkn         string
	adminChatID int64
}

// Token returns the bot token
func (b BotConfig) Token() string {
	return b.tkn
}

// GetAdminChatID returns admin chat id
func (b BotConfig) GetAdminChatID() int64 {
	return b.adminChatID
}
