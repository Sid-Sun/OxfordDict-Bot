package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"
	"go.uber.org/zap"
	"time"
)

// RedisSoreImpl implements Store with redis
type RedisSoreImpl struct {
	rc     *redis.Client
	logger *zap.Logger
}

func NewRedisStore(rdc *redis.Client, lgr *zap.Logger) Store {
	return &RedisSoreImpl{
		rc:     rdc,
		logger: lgr,
	}
}

// Get returns a db Data instance corresponding to id
func (i *RedisSoreImpl) Get(id string) api.Response {
	var d api.Response
	res, err := i.rc.Get(context.Background(), id).Result()
	if err != nil {
		if err != redis.Nil {
			i.logger.Error(fmt.Sprintf("[Store] [Instance] [Get] [Get] %v", err))
		}
		return d
	}
	err = json.Unmarshal([]byte(res), &d)
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Store] [Instance] [Get] [Unmarshal] %v", err))
	}
	return d
}

// Put unconditionally sets db record of id to provided data
func (i *RedisSoreImpl) Put(id string, data api.Response) {
	d, err := json.Marshal(data)
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Store] [Instance] [Put] [Marshal] %v", err))
		return
	}
	err = i.rc.Set(context.Background(), id, d, 24*time.Hour).Err()
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Store] [Instance] [Put] [Set] %v", err))
	}
}
