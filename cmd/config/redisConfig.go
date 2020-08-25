package config

import "fmt"

// RedisConfig definges configuration for redis server
type RedisConfig struct {
	host     string
	port     int
	password string
	db       int
}

// Address returns the Address of redis server
func (r RedisConfig) Address() string {
	return fmt.Sprintf("%s:%d", r.host, r.port)
}

// Password returns the password for connecting to redis server 
func (r RedisConfig) Password() string {
	return r.password
}

// DB returns the DB id for redis
func (r RedisConfig) DB() int {
	return r.db
}
