package main

import (
	"fmt"
	"sync"

	"github.com/Path94/atoms"
)

func main() {
	var wg sync.WaitGroup
	c := NewCounter()
	wg.Add(3)

	go func() {
		c.Increment("foo")
		wg.Done()
	}()

	go func() {
		c.Increment("foo")
		wg.Done()
	}()

	go func() {
		c.Increment("foo")
		wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Output will be 3
	fmt.Println(c.Get("foo"))
}

// NewCounter will return a new counter
func NewCounter() *Counter {
	var c Counter
	// Initialize our internal map
	c.cm = make(map[string]uint64)
	return &c
}

// Counter manages a counter
type Counter struct {
	mux atoms.RWMux
	cm  map[string]uint64
}

// Increment will increment a given key
func (c *Counter) Increment(key string) {
	c.mux.Update(func() {
		c.cm[key]++
	})
}

// Get will get the counter value for a given key
func (c *Counter) Get(key string) (n uint64) {
	c.mux.Read(func() {
		n = c.cm[key]
	})

	return
}
