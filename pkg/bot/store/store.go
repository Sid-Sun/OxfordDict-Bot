package store

import "github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"

// Store wraps InstanceInterface
type Store interface {
	Get(string) api.Response
	Put(string, api.Response)
}
