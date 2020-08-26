package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config contains all the necessary configurations
type Config struct {
	Bot         BotConfig
	environment string
	API         APIConfig
}

// GetEnv returns the current development environment
func (c Config) GetEnv() string {
	return c.environment
}

// Load reads all config from env to config
func Load() Config {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	return Config{
		environment: viper.GetString("APP_ENV"),
		Bot: BotConfig{
			tkn:         viper.GetString("API_TOKEN"),
			adminChatID: viper.GetInt64("ADMIN_CHAT_ID"),
		},
		API: NewAPIConfig(strings.Split(viper.GetString("APP_IDS"), ";"), strings.Split(viper.GetString("APP_KEYS"), ";")),
	}
}
