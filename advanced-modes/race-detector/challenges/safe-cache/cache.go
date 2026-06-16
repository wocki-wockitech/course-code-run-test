package main

// Cache is a thread-safe string key-value store.
// TODO: add fields and implement methods without data races.
type Cache struct {
	data map[string]string
}

// NewCache creates a new Cache instance.
func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

// Set stores a value by key.
func (c *Cache) Set(key, value string) {
	c.data[key] = value
}

// Get retrieves a value by key. Returns the value and whether it exists.
func (c *Cache) Get(key string) (string, bool) {
	v, ok := c.data[key]
	return v, ok
}
