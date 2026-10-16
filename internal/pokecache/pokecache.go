package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Mu       sync.Mutex
	Cache    map[string]CacheEntry
	Interval time.Duration
}

// NewCache creates a new cache with a specified cleanup interval
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Cache:    make(map[string]CacheEntry),
		Interval: interval,
	}

	// Start the background reaping loop
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	if entry, ok := c.Cache[key]; ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

// reapLoop runs periodically to delete entries older than the interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Mu.Lock()
		now := time.Now()
		for key, entry := range c.Cache {
			if now.Sub(entry.createdAt) > interval {
				delete(c.Cache, key)
			}
		}
		c.Mu.Unlock()
	}
}
