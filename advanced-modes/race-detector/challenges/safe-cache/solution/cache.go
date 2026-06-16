package main

import "sync"

// Cache is a thread-safe string key-value store.
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewCache creates a new Cache instance.
func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

// Set stores a value by key.
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get retrieves a value by key. Returns the value and whether it exists.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}
