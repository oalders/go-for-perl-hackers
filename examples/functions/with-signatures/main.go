// package main prints the lesser of two integers
package main

import "fmt"

func main() {
	fmt.Println(min(3, 4))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
