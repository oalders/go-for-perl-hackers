// Package main demonstrates how to use "waitgroup". Examples are courtesy of
// Claude 3.5 Haiku.

// Essentially waitgroup provides a way to run concurrent functions, but wait
// until the have all run to completion before continuing on.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to simulate fetching data from a website
func fetchWebsite(url string, wg *sync.WaitGroup, results chan string) {
	// Ensure WaitGroup is decremented when the goroutine completes
	defer wg.Done()

	// Simulate network delay
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

	// Simulate fetching data
	results <- fmt.Sprintf("Data from %s", url)
}

func main() {
	// List of websites to fetch
	websites := []string{
		"https://example.com",
		"https://another-site.org",
		"https://sample-website.net",
	}

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Channel to collect results
	results := make(chan string, len(websites))

	// Launch a goroutine for each website
	for _, website := range websites {
		// Increment the WaitGroup counter before starting the goroutine
		wg.Add(1)

		// Start goroutine to fetch website
		go fetchWebsite(website, &wg, results)
	}

	// Wait for all goroutines to complete
	// This prevents the main function from exiting before goroutines finish
	wg.Wait()

	// Close the results channel
	close(results)

	// Collect and print results
	fmt.Println("Fetched results:")
	for result := range results {
		fmt.Println(result)
	}
}
