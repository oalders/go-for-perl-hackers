// Package main demonstrates how to recover from a panic.
package main

import "fmt"

func main() {
	{
		result, err := riskyMath(2, 0)
		fmt.Printf("result: %d err: %s\n", result, err)
	}
	{
		result, err := riskyMath(1, 2)
		fmt.Printf("result: %d err: %s\n", result, err)
	}
}

func riskyMath(a, b int) (result int, err error) {
	// Only use recover in deferred functions. It only works when called
	// directly from a deferred function.
	defer func() {
		if r := recover(); r != nil {
			result = 0
			err = fmt.Errorf("division error: %v", r)
		}
	}()
	return a / b, nil
}
