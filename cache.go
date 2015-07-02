package gocache

import "time"

// Cacher is an interface to some caching engine.
type Cacher interface {
	Put(key string, value interface{}) error
	PutPersist(key string, value interface{}) error
	PutTtl(key string, value interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
}
