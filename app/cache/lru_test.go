package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		lruCache := NewLRUCache(1*time.Second, 10, 1)

		result1, err1 := lruCache.Get("randomkey")
		assert.Equal(t, nil, result1)

		assert.NoError(t, err1)

		err2 := lruCache.Set("randomkey", "randomvalue")
		assert.NoError(t, err2)

		resultFromCache, err3 := lruCache.Get("randomkey")
		assert.Equal(t, resultFromCache, "randomvalue")
		assert.NoError(t, err3)

		result2, err4 := lruCache.Get("unexistentkey")
		assert.Equal(t, nil, result2)
		assert.NoError(t, err4)
	})
}
