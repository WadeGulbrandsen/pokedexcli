package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	lock    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		lock:    &sync.RWMutex{},
	}
	go cache.reaper(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.entries[key] = cacheEntry{createdAt: time.Now().UTC(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reaper(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for k, entry := range c.entries {
		if time.Since(entry.createdAt) > interval {
			delete(c.entries, k)
		}
	}
}
