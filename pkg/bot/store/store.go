package store

// Store wraps InstanceInterface
type Store struct {
	InstanceInterface
}

// NewStore creates a new Store instance
func NewStore(instanceInterface InstanceInterface) Store {
	return Store{instanceInterface}
}
