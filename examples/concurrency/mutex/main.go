// Package main demonstrates how to use a "mutex" lock (mutually exclusive
// lock). Examples are courtesy of Claude 3.5 Haiku.
package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	// Mutex to protect the counter
	mu sync.Mutex

	// Counter value
	value int
}

// Increment safely increments the counter
func (c *SafeCounter) Increment() {
	// Lock the mutex before modifying the counter
	c.mu.Lock()

	// Ensure the mutex is always unlocked using defer
	// This happens even if the function panics
	defer c.mu.Unlock()

	c.value++
}

// Value returns the current counter value
func (c *SafeCounter) Value() int {
	// Lock mutex when reading to prevent race conditions
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func main() {
	// Create a safe counter
	counter := &SafeCounter{}

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Number of goroutines to spawn
	numGoroutines := 100

	// Increment counter concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Each goroutine increments counter multiple times
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print final counter value
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
