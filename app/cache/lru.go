package cache

import (
	"fmt"
	"time"

	"github.com/karlseguin/ccache"
	"github.com/rs/zerolog/log"
)

type lruCache struct {
	name  string
	cache *ccache.Cache
	ttl   time.Duration
}

func NewLRUCache(ttl time.Duration, maxSize int64, itemsToPrune uint32) Cache {
	minutes := time.Now().Minute()
	hour := time.Now().Hour()
	return &lruCache{
		name:  string(fmt.Sprintf("cache-%d-%d", hour, minutes)),
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
	log.Debug().Msgf("%v is the key, %v is the value and %v is the the ttl", key, value, lc.ttl)
	return nil
}

func (lc *lruCache) GetName() string {
	return lc.name
}
