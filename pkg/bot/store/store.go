package store

import "github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"

// Store wraps InstanceInterface
type Store interface {
	Get(string) api.Response
	Put(string, api.Response)
	DoesUserExist(ChatID int64) bool
	InsertUser(ChatID int64)
	IncrementQueryCount(ChatID int64)
}
