// Package main is a basic demonstrations of channels
// Stolen from https://go.dev/tour/concurrency/2
package main

import "fmt"

// The "<-" indicates the direction in which data flows.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8}
	b := []int{-9, 4, 0}

	pipeline := make(chan int)

	// The goroutines are run concurrently, but the channel receive is blocking.
	// This means that we can take advantage of parallelism when running
	// computations, but we can also ensure a proper order of operations
	// because of the blocking on the channel receive.
	go sum(a, pipeline) // these operations are concurrent
	go sum(b, pipeline) // these operations are concurrent

	// this could also have been written as
	// x := <-c
	// y := <-c
	x, y := <-pipeline, <-pipeline // receive from c. this is blocking

	fmt.Println(x, y, x+y)
}
