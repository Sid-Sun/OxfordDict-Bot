package config

import (
	"os"
)

// Config contains all the necessary configurations
type Config struct {
	Bot         BotConfig
	environment string
}

// GetEnv returns the current development environment
func (c Config) GetEnv() string {
	return c.environment
}

// Load reads all config from env to config
func Load() Config {
	return Config{
		environment: os.Getenv("APP_ENV"),
		Bot: BotConfig{
			tkn: os.Getenv("API_TOKEN"),
		},
	}
}
