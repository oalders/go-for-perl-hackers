// Package main demonstrates how to implement the Stringer interface defined in
// the fmt package.

package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) String() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	person := Person{"Olaf", "Alders"}
	fmt.Println(person)
}
