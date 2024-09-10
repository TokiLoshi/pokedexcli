package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	value []byte
}



func NewCache(interval time.Duration) Cache {
	fmt.Println("Creating a new cache", interval)
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		value: value,
	}
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, exists := c.cache[key]
	return entry.value, exists
}

func (c *Cache)reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
	fmt.Println("Cleaning up in ReapLoop")
	

}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, value := range c.cache {
			if value.createdAt.Before(now.Add(-last)){
				delete(c.cache, key)
			}
	}
}