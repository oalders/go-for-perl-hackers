package main
// Package main demonstrates how to use "select" with a default branch.
// Examples are courtesy of Claude 3.5 Haiku.

import (
    "fmt"
    "time"
)

func main() {
    // Create channels
    ch1 := make(chan string)
    ch2 := make(chan int)

    // Demonstrate a select with default branch
    select {
    case msg1 := <-ch1:
        fmt.Println("Received from ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from ch2:", msg2)
    default:
        fmt.Println("No channel is ready - continuing without blocking")
    }

    // Real-world example: non-blocking channel check
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "delayed message"
    }()

    // Non-blocking select that continues if no message is available
    for i := 0; i < 8; i++ {
        select {
        case msg := <-ch1:
            fmt.Println("Received:", msg)
        default:
            fmt.Println("No message available, continuing...")
        }
        time.Sleep(500 * time.Millisecond)
    }
}
