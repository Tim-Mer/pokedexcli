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
	memory map[string]CacheEntry
	mutex  sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	//creates a new cache with a refresh of interval
	defaultCache := Cache{memory: make(map[string]CacheEntry)}
	go defaultCache.reapLoop(interval)
	return &defaultCache
}

func (c *Cache) Add(key string, val []byte) {
	//Adds an entry to the cache
	c.mutex.Lock()
	tmpEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.memory[key] = tmpEntry
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	//Gets an entry from the cache and returns it along with a bool that says if it was found or not
	c.mutex.Lock()
	val, ok := c.memory[key]
	c.mutex.Unlock()
	if ok { // if entry exists
		return val.val, true
	} // else return false
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	//Deletes cache entries that are old
	ticker := time.NewTicker(interval)

	defer ticker.Stop()
	for t := range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.memory {
			if t.Sub(entry.createdAt) > interval {
				// delete the entry
				delete(c.memory, key)
			}
		}
		c.mutex.Unlock()
	}
}
