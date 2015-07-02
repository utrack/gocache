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
func (c *GoCache) New(exptime time.Duration, ptime time.Duration) {
	c.ce = cache.New(exptime, ptime)
}

func (c *GoCache) Put(key string, value interface{}) {
	c.ce.Set(key, value, cache.DefaultExpiration)
}

func (c *GoCache) PutPersist(key string, value interface{}) {
	c.ce.Set(key, value, cache.NoExpiration)
}

func (c *GoCache) PutTtl(key string, value interface{}, ttl time.Duration) {
	c.ce.Set(key, value, ttl)
}

func (c *GoCache) Get(key string) (interface{}, error) {
	ret, _ := c.ce.Get(key)
	return ret, nil
}
