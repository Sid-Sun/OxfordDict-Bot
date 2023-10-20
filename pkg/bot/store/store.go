package store

// Store wraps Redis
type Store struct {
	Redis RedisService
}

// NewStore creates a new Store instance
func NewStore(Redis RedisService) Store {
	return Store{
		Redis: Redis,
	}
}
