package gocache

import "time"

// Cacher is an interface to some caching engine.
type Cacher interface {
	// Put saves the data at specified key.
	Put(key string, value interface{}, cacheTtl time.Duration) error
	// Get retrieves data from the cache and boolean value
	// indicating if the value was found.
	Get(key string) (interface{}, bool)
}

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)
