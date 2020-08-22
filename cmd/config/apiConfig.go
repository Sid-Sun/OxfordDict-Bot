package config

// APIConfig defines config for API
type APIConfig struct {
	id  string
	key string
}

// GetID returns app ID for API
func (a APIConfig) GetID() string {
	return a.id
}

// GetKey returns key for API calls
func (a APIConfig) GetKey() string {
	return a.key
}
