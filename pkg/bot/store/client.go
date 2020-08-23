package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"go.uber.org/zap"
)

// ClientInterface defines methods for client
type ClientInterface interface {
	GetClient() (*redis.Client, error)
}

// Client implements ClientInterface
type Client struct {
	config config.RedisConfig
	logger *zap.Logger
}

// NewClient returns a new instance of Client with necessary config
func NewClient(lgr *zap.Logger, cfg config.RedisConfig) ClientInterface {
	return Client{
		config: cfg,
		logger: lgr,
	}
}

// GetClient creates and returns a new Redis client
func (c Client) GetClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.config.Address(),
		Password: c.config.Password(),
		DB:       c.config.DB(),
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return &redis.Client{}, err
	}
	return rdb, nil
}
