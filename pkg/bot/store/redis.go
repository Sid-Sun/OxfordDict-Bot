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

// InstanceInterface defines a db instance interface
type InstanceInterface interface {
	Get(string) api.Response
	Put(string, api.Response)
}

// NewInstance creates a new instance for db
func NewInstance(rdb *redis.Client, logger *zap.Logger) InstanceInterface {
	return Instance{logger: logger, rc: rdb}
}

// Instance implements InstanceInterface with map
type Instance struct {
	rc     *redis.Client
	logger *zap.Logger
}

// Get returns a db Data instance corresponding to id
func (i Instance) Get(id string) api.Response {
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
func (i Instance) Put(id string, data api.Response) {
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
