package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	res := Cache{entry: map[string]cacheEntry{}, mu: &sync.Mutex{}}

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			res.reapLoop(interval)
		}
	}()

	return res
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{val: val, createdAt: time.Now()}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	res, ok := c.entry[key]
	return res.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entry {
		if time.Since(entry.createdAt) > interval {
			delete(c.entry, key)
		}
	}
}
