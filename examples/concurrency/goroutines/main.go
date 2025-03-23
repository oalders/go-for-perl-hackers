// Package main demonstrates basic use of goroutines
// Stolen from https://go.dev/tour/concurrency/1
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s (%d)\n", s, i)
	}
}

func main() {
	go say("world")
	say("hello")
}
