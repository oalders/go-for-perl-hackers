// Package main demonstrates how to use context.WithTimeout via net/http.
// Examples are courtesy of Claude 3.5 Haiku.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Create a context that will time out after 3 seconds
	//
	// Background returns a non-nil, empty [Context]. It is never canceled, has no
	// values, and has no deadline. It is typically used by the main function,
	// initialization, and tests, and as the top-level Context for incoming
	// requests.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Always remember to call cancel to release resources

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://olafalders.com", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Check if the error is due to context timeout
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out after 3 seconds")
		} else {
			fmt.Printf("Error making request: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response length: %d bytes\n", len(body))
	fmt.Println("First 100 bytes of response:")
	if len(body) > 100 {
		fmt.Println(string(body[:100]) + "...")
	} else {
		fmt.Println(string(body))
	}
}
