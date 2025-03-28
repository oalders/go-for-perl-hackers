// Package main demonstrates how to use "once". Examples are courtesy of
// Claude 3.5 Haiku.
//
// Initialization Guarantee
//
// The function passed to once.Do() is executed exactly one time
// Subsequent calls are no-ops
// Thread-safe initialization mechanism
package main

import (
    "fmt"
    "sync"
)

// Expensive resource that we want to initialize only once
type ExpensiveResource struct {
    data string
}

// Singleton-like resource manager
type ResourceManager struct {
    resource *ExpensiveResource
    once     sync.Once
}

func (rm *ResourceManager) GetResource() *ExpensiveResource {
    // This initialization will happen only ONCE, no matter how many times it's called
    rm.once.Do(func() {
        fmt.Println("Initializing expensive resource...")
        rm.resource = &ExpensiveResource{
            data: "Initialized Resource Data",
        }
    })
    return rm.resource
}

func main() {
    manager := &ResourceManager{}

    // Concurrent calls to demonstrate Once behavior
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            resource := manager.GetResource()
            fmt.Printf("Goroutine %d: %v\n", id, resource.data)
        }(i)
    }

    wg.Wait()
}
