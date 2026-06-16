package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestCacheBasic(t *testing.T) {
	c := NewCache()
	c.Set("hello", "world")
	v, ok := c.Get("hello")
	if !ok || v != "world" {
		t.Errorf("Get(hello) = %q, %v; want world, true", v, ok)
	}
	_, ok = c.Get("missing")
	if ok {
		t.Error("Get(missing) should return false")
	}
}

func TestCacheConcurrentWrites(t *testing.T) {
	c := NewCache()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Set(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d", i))
		}(i)
	}
	wg.Wait()

	for i := 0; i < 100; i++ {
		v, ok := c.Get(fmt.Sprintf("key-%d", i))
		if !ok || v != fmt.Sprintf("val-%d", i) {
			t.Errorf("key-%d: got %q, %v", i, v, ok)
		}
	}
}

func TestCacheConcurrentReadWrite(t *testing.T) {
	c := NewCache()
	c.Set("shared", "initial")

	var wg sync.WaitGroup
	// Writers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Set("shared", fmt.Sprintf("v%d", i))
		}(i)
	}
	// Readers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("shared")
		}()
	}
	wg.Wait()

	// Just verify it doesn't panic and race detector is happy.
	_, ok := c.Get("shared")
	if !ok {
		t.Error("shared key should exist")
	}
}
