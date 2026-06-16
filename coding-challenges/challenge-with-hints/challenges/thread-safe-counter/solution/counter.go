package main

import "sync"

// Counter is a thread-safe integer counter.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter by 1.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Dec decrements the counter by 1.
func (c *Counter) Dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

// Value returns the current counter value.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
