package config

import (
	"github.com/spf13/viper"
)

// Config contains all the necessary configurations
type Config struct {
	Bot         BotConfig
	environment string
	APIKey      string
	Redis       RedisConfig
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
		APIKey: viper.GetString("API_KEY"),
		Redis: RedisConfig{
			host:     viper.GetString("REDIS_HOST"),
			port:     viper.GetInt("REDIS_PORT"),
			password: viper.GetString("REDIS_PASS"),
			db:       viper.GetInt("REDIS_DB"),
			ssl:      viper.GetBool("REDIS_SSL"),
		},
	}
}
