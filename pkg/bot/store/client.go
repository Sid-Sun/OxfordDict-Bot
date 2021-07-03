package store

import (
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"go.uber.org/zap"
)

// NewRedisClient returns a new instance of Client with necessary config
func NewRedisClient(cfg config.RedisConfig, lgr *zap.Logger) (*redis.Client, error) {
	var t *tls.Config
	if cfg.SSL() {
		t = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:      cfg.Address(),
		Password:  cfg.Password(),
		DB:        cfg.DB(),
		TLSConfig: t,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return &redis.Client{}, err
	}
	return rdb, nil
}
