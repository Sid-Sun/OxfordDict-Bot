package config

import "fmt"

type DBConfig struct {
	port     int
	server   string
	user     string
	password string
	database string
}

func (d DBConfig) GetConn() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		d.server, d.user, d.password, d.port, d.database)
}
