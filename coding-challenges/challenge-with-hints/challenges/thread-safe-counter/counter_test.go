package main

import (
	"sync"
	"testing"
)

func TestCounterBasic(t *testing.T) {
	c := &Counter{}
	c.Inc()
	c.Inc()
	c.Inc()
	c.Dec()
	if v := c.Value(); v != 2 {
		t.Errorf("expected 2, got %d", v)
	}
}

func TestCounterConcurrent(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Dec()
		}()
	}
	wg.Wait()

	if v := c.Value(); v != 500 {
		t.Errorf("expected 500 after concurrent ops, got %d", v)
	}
}

func TestCounterZero(t *testing.T) {
	c := &Counter{}
	if v := c.Value(); v != 0 {
		t.Errorf("new counter should be 0, got %d", v)
	}
}
