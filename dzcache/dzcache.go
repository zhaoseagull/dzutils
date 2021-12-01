package dzcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type DzCacheStruct struct {
	DefaultExipre time.Duration
	PurgeExipre   time.Duration
	DzCache       *cache.Cache
}

func NewDzCache(defaultExipre, purgeExipre time.Duration) *DzCacheStruct {
	return &DzCacheStruct{
		DefaultExipre: defaultExipre,
		PurgeExipre:   purgeExipre,
		DzCache:       cache.New(defaultExipre, purgeExipre),
	}
}

func (c *DzCacheStruct) Set(key string, value interface{}, expire time.Duration) {
	c.DzCache.Set(key, value, expire)
}
func (c *DzCacheStruct) Get(key string) interface{} {
	if v, ok := c.DzCache.Get(key); ok {
		return v
	} else {
		return nil
	}
}
