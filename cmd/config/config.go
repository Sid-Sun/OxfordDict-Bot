package config

import (
	"os"
	"strconv"
)

// Config contains all the necessary configurations
type Config struct {
	Bot         BotConfig
	environment string
	API         APIConfig
	Redis       RedisConfig
}

// GetEnv returns the current development environment
func (c Config) GetEnv() string {
	return c.environment
}

// Load reads all config from env to config
func Load() Config {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(err)
	}
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}

	return Config{
		environment: os.Getenv("APP_ENV"),
		Bot: BotConfig{
			tkn:         os.Getenv("API_TOKEN"),
			adminChatID: os.Getenv("ADMIN_CHAT_ID"),
		},
		API: APIConfig{
			id:  os.Getenv("APP_ID"),
			key: os.Getenv("APP_KEY"),
		},
		Redis: RedisConfig{
			host:     os.Getenv("REDIS_HOST"),
			port:     port,
			password: os.Getenv("REDIS_PASS"),
			db:       db,
		},
	}
}
