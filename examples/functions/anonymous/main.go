// package main demonstrates how to create anonymous functions. They are
// anonymous because they are not named. The function can be assigned to a
// value with a name, but the function itself has no name.
package main

import "fmt"

func main() {
	// assign the function directly to a variable
	f := func() {
		fmt.Println("anonymous function")
	}
	f()

	creation := factory()
	fmt.Println(creation())
}

// A function which creates and returns a different function. Note the
// signature of the factory() function.
func factory() func() int {
	return func() int {
		return 5
	}
}
