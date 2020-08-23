package config

import "fmt"

type RedisConfig struct {
	host     string
	port     int
	password string
	db       int
}

func (r RedisConfig) Address() string {
	return fmt.Sprintf("%s:%d", r.host, r.port)
}

func (r RedisConfig) Password() string {
	return r.password
}

func (r RedisConfig) DB() int {
	return r.db
}
