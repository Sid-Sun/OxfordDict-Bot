package config

import (
	"errors"

	queue "github.com/nsnikhil/go-datastructures/queue"
)

// Node contains the id and key for API
type Node struct {
	id  string
	key string
}

// GetID returns app ID for API
func (c Node) GetID() string {
	return c.id
}

// GetKey returns key for API calls
func (c Node) GetKey() string {
	return c.key
}

// APIConfig defines config for API
type APIConfig struct {
	lqueue *queue.LinkedQueue
}

// GetConfig returns a config instance
func (c *APIConfig) GetConfig() (Node, error) {
	node, error := c.lqueue.Remove()
	c.lqueue.Add(node)
	return node.(Node), error
}

// NewAPIConfig creates a new APIConfig and list
func NewAPIConfig(ids, keys []string) APIConfig {
	if len(ids) != len(keys) || len(ids) == 0 {
		panic(errors.New("length of ids and keys should be equal and not 0"))
	}

	lq, err := queue.NewLinkedQueue()
	if err != nil {
		panic(err)
	}

	for i, id := range ids {
		lq.Add(Node{
			id:  id,
			key: keys[i],
		})
	}

	return APIConfig{
		lqueue: lq,
	}
}
