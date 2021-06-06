package cache

import (
	"time"
)

// MockedCache is an empty struct used just for testing
type mockedCache struct {
	cache *Cache
	ttl   time.Duration
}

// Get returns an empty byte array and a nil error
func (m *mockedCache) Get(key string) (interface{}, error) {
	return nil, nil
}

// Set returns a nil error
func (m *mockedCache) Set(key string, value interface{}) error {
	return nil
}

// NewMockedCache creates a mocked cache just for testing
func NewMockedCache(ttl time.Duration, maxSize int64, itemsToPrune uint32) Cache {
	return &mockedCache{
		cache: nil,
		ttl:   ttl,
	}
}
