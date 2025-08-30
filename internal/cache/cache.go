package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value     any
	ExpiresAt *time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]*CacheItem
}

func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]*CacheItem),
	}

	// Start cleanup
	go c.startCleanup()

	return c
}

func (c *Cache) Set(key string, value any) error {
	return c.SetWithTTL(key, value, 0)
}

func (c *Cache) SetWithTTL(key string, value any, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := &CacheItem{
		Value: value,
	}

	if ttl > 0 {
		expiresAt := time.Now().Add(ttl)
		item.ExpiresAt = &expiresAt
	}

	c.items[key] = item
	return nil
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	// check is expired
	if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
		// remove expires item (need write lock)
		c.mu.RUnlock()
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		c.mu.RLock()
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.items[key]
	if exists {
		delete(c.items, key)
	}

	return exists
}

// start clean up cicle
func (c *Cache) startCleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.cleanup()
	}
}

// clena all expired items
func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, item := range c.items {
		if item.ExpiresAt != nil && now.After(*item.ExpiresAt) {
			delete(c.items, key)
		}
	}
}
