package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
	ticker   *time.Ticker
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{cache: map[string]cacheEntry{}, mu: sync.Mutex{}, interval: interval, ticker: time.NewTicker(interval)}

	go cache.reapLoop()

	return &cache
}

// use sync.Mutex as appropiate
// adds a new entry to the cache. It should take a key(a string) and a val( a []byte)
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
}

// Create a cache.Get() method that gets an entry from the cache.
// It should take a key(a string) and return a []byte and a bool.
// The bool should be true if the entry was found and false if it wasn't
// use sync.Mutex as appropiate
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce, ok := c.cache[key]
	return ce.val, ok
}

// Create a cache.reapLoop() method that is called when the cache is created(by the NewCache function).
// Each time an interval (the time.Duration passed to NewCache) passes it should remove any entries that are older than the interval.
// This makes sure that the cache doesn't grow too large over time. For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.
// I used a time.Ticker to make this happen. If you want addition help, see the Tips section below
// use sync.Mutex as appropiate
func (c *Cache) reapLoop() {
	for t := range c.ticker.C {
		for k, e := range c.cache {
			delta := t.Sub(e.createdAt)
			if delta >= c.interval {
				delete(c.cache, k)
			}
		}
	}
}
