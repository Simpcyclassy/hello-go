package cache

import (
	"time"

	"github.com/karlseguin/ccache"
)

type lruCache struct {
	cache *ccache.Cache
	ttl   time.Duration
}

func NewLRUCache(ttl time.Duration, maxSize int64, itemsToPrune uint32) Cache {
	return &lruCache{
		cache: ccache.New(ccache.Configure().MaxSize(maxSize).ItemsToPrune(itemsToPrune)),
		ttl:   ttl,
	}
}

func (lc *lruCache) Get(key string) (interface{}, error) {
	result := lc.cache.Get(key)
	if result == nil || result.Expired() {
		return nil, nil
	}
	return result.Value(), nil
}

func (lc *lruCache) Set(key string, value interface{}) error {
	lc.cache.Set(key, value, lc.ttl)
	return nil
}
