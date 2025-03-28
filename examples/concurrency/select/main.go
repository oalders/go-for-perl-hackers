// Package main demonstrates how to use "select". Examples are courtesy of
// Claude 3.5 Haiku.
package main

import (
	"fmt"
	"time"
)

// Scenario 1: Handling Multiple Channel Operations
// In this scenario the first channel to return data will be printed.
func multiChannelCommunication() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Message from channel 1"
	}()

	go func() {
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	}
}

// Scenario 2: Timeout Handling
// time.After() returns a <-chan Time
func timeoutExample() {
	ch := make(chan string)

	go func() {
		// Simulate a long-running operation
		time.Sleep(2 * time.Second)
		ch <- "Operation completed"
	}()

	fmt.Println("It is now: " + time.Now().Format("15-04-05"))
	select {
	case result := <-ch:
		fmt.Println("Received result:", result)
	case timedOutAt := <-time.After(2 * time.Second):
		fmt.Println("Operation timed out at " + timedOutAt.Format(("15-04-05")))
	}
}

// Scenario 3: Non-blocking Channel Operations

// The `1` in `ch := make(chan int, 1)` specifies the capacity of the channel,
// making it a **buffered channel** with a buffer size of 1.
//
// A buffered channel can hold a limited number of values without requiring a
// receiver to be ready to consume them immediately. In this case, the channel can
// store up to one `int` value. If the channel is full (i.e., it already contains
// one value), any attempt to send another value to the channel will block until a
// receiver retrieves the existing value. Conversely, if the channel is empty, a
// receive operation will block until a value is sent to the channel.
//
// This small buffer size is useful for scenarios where you want to allow a single
// value to be sent without blocking the sender, but still enforce synchronization
// when the buffer is full.

func nonBlockingReceive() {
	ch := make(chan int, 1)

	select {
	case val := <-ch:
		fmt.Println("Received value:", val)
	default:
		fmt.Println("No value available")
	}

	// Add a value to the buffered channel
	ch <- 42

	select {
	case val := <-ch:
		fmt.Println("Received value:", val)
	default:
		fmt.Println("No value available")
	}
}

// Scenario 4: Coordinating Multiple Goroutines
func coordinateGoroutines() {
	done := make(chan bool)
	data := make(chan int)

	go func() {
		// Simulating data processing
		time.Sleep(500 * time.Millisecond)
		data <- 100
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Processing completed")
	case <-data:
		fmt.Println("Data received")
	case <-time.After(1 * time.Second):
		fmt.Println("Operation took too long")
	}
}

func main() {
	fmt.Println("Scenario 1: Multi-Channel Communication")
	multiChannelCommunication()

	fmt.Println("\nScenario 2: Timeout Handling")
	timeoutExample()

	fmt.Println("\nScenario 3: Non-blocking Receive")
	nonBlockingReceive()

	fmt.Println("\nScenario 4: Coordinating Goroutines")
	coordinateGoroutines()
}
