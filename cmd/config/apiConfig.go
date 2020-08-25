package config

import (
	"errors"

	list "github.com/nsnikhil/go-datastructures/list"
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
	current      Node
	currentIndex int
	total        int
	list         *list.ArrayList
}

// GetConfig returns a config instance
func (c *APIConfig) GetConfig() Node {
	cfg := c.current

	c.currentIndex = c.currentIndex + 1
	if c.currentIndex >= c.total {
		c.currentIndex = 0
	}

	c.current = c.list.Get(c.currentIndex).(Node)

	return cfg
}

// NewAPIConfig creates a new APIConfig and list
func NewAPIConfig(ids, keys []string) APIConfig {
	if len(ids) != len(keys) || len(ids) == 0 {
		panic(errors.New("length of ids and keys should be equal and not 0"))
	}

	al, err := list.NewArrayList()
	if err != nil {
		panic(err)
	}

	for i, id := range ids {
		al.Add(Node{
			id:  id,
			key: keys[i],
		})
	}

	return APIConfig{
		current:      al.Get(0).(Node),
		currentIndex: 0,
		total:        len(ids),
		list:         al,
	}
}
