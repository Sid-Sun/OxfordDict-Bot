package store

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
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

func InitMSSQLClient(cfg *config.DBConfig, lgr *zap.Logger) (*sql.DB, error) {
	var err error

	// Create connection pool
	db, err := sql.Open("sqlserver", cfg.GetConn())
	if err != nil {
		lgr.Fatal(fmt.Sprintf("[Store] [InitMSSQLClient] [Open] %s", err.Error()))
		return nil, err
	}

	err = db.PingContext(context.Background())
	if err != nil {
		lgr.Fatal(fmt.Sprintf("[Store] [InitMSSQLClient] [PingContext] %s", err.Error()))
		return nil, err
	}

	return db, nil
}
