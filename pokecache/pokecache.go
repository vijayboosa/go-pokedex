package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex    sync.Mutex
	items    map[string]CacheEntry
	interval time.Duration
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	item, ok := c.items[key]
	if !ok {
		return nil, ok
	}
	return item.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		currentTime := time.Now().Add(c.interval)
		c.mutex.Lock()
		for key, value := range c.items {
			cacheEndTime := value.createdAt.Add(c.interval)
			if currentTime.Before(cacheEndTime) {
				continue
			}
			delete(c.items, key)
		}
		c.mutex.Unlock()

	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		items:    map[string]CacheEntry{},
	}

	go cache.reapLoop()

	return cache
}
