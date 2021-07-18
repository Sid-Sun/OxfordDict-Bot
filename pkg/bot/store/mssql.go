package store

import (
	"database/sql"
	"fmt"

	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"
	"go.uber.org/zap"
)

type MSSQLRedisStore struct {
	redisStore RedisSoreImpl
	sql        *sql.DB
	lgr        *zap.Logger
}

func NewMSSQLRedisStore(rds RedisSoreImpl, sql *sql.DB, lgr *zap.Logger) Store {
	return &MSSQLRedisStore{
		redisStore: rds,
		sql: sql,
		lgr: lgr,
	}
}

func (m *MSSQLRedisStore) Get(id string) api.Response {
	return m.redisStore.Get(id)
}

func (m *MSSQLRedisStore) Put(id string, data api.Response) {
	m.redisStore.Put(id, data)
}

func (m *MSSQLRedisStore) DoesUserExist(ChatID int64) bool {
	query := `SELECT user_id FROM users WHERE user_id=@userID`

	row := m.sql.QueryRow(query, sql.Named("userID", ChatID))
	var id int64
	err := row.Scan(&id)
	if err != nil {
		m.lgr.Error(fmt.Sprintf("[Store] [MSSQLRedisStore] [DoesUserExist] [Scan] %s", err.Error()))
	}

	return id != 0
}

func (m *MSSQLRedisStore) InsertUser(ChatID int64) {
	query := `INSERT INTO users (user_id, query_count) VALUES (@userID, @queryCount)`

	_, err := m.sql.Exec(query, sql.Named("userID", ChatID), sql.Named("queryCount", 1))
	if err != nil {
		m.lgr.Error(fmt.Sprintf("[Store] [MSSQLRedisStore] [InsertUser] [Exec] %s", err.Error()))
	}
}

func (m *MSSQLRedisStore) IncrementQueryCount(ChatID int64) {
	query := `UPDATE users SET query_count = query_count + 1 WHERE user_id=@userID`

	_, err := m.sql.Exec(query, sql.Named("userID", ChatID))
	if err != nil {
		m.lgr.Error(fmt.Sprintf("[Store] [MSSQLRedisStore] [IncrementQueryCount] [Exec] %s", err.Error()))
	}
}
