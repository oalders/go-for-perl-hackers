// Package main demonstrates a simple use of interface
// https://go.dev/tour/methods/9
package main

import "fmt"

// Greeter defines an interface, which has the sole requirement that the struct
// implement a method named Waves() which returns a string.
type Greeter interface {
	Waves() string
}

// The Usher type just contains a name.
type Usher struct {
	Name string
}

// Usher also implements the Waves() method
func (u Usher) Waves() string {
	return "o/"
}

// hello() expects something which implements the Greeter interface.
func hello(g Greeter) {
	fmt.Println(g.Waves())
}

func main() {
	employee := Usher{Name: "Mario"}

	// Because Usher implements a Waves() method, we can now pass it to the
	// hello() function.
	hello(employee)
}
