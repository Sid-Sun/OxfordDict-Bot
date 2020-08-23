package store

import (
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"
	"go.uber.org/zap"
)

// InstanceInterface defines a db instance interface
type InstanceInterface interface {
	Get(string) api.Response
	Put(string, api.Response)
}

// NewInstance creates a new instance for db
func NewInstance(logger *zap.Logger) InstanceInterface {
	return Instance{logger: logger, data: make(map[string]api.Response)}
}

// Instance implements InstanceInterface with map
type Instance struct {
	data   map[string]api.Response
	logger *zap.Logger
}

// Get returns a db Data instance corresponding to id
func (i Instance) Get(id string) api.Response {
	return i.data[id]
}

// Put unconditionally sets db record of id to provided data
func (i Instance) Put(id string, data api.Response) {
	i.data[id] = data
}
