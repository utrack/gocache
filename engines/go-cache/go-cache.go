package cacheGoCache

import (
	"time"

	"github.com/pmylund/go-cache"
)

type GoCache struct {
	ce *cache.Cache
}

// New initiates go-cache cacher with supplied default expiration
// time and purge time.
func New(exptime time.Duration, ptime time.Duration) *GoCache {
	ret := new(GoCache)
	ret.ce = cache.New(exptime, ptime)
	return ret
}

func (c *GoCache) Put(key string, value interface{}, ttl time.Duration) {
	c.ce.Set(key, value, ttl)
}

func (c *GoCache) Get(key string) (interface{}, bool) {
	return c.ce.Get(key)
}
