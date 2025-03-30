// Package main demonstrates the use of type assertions.
// A type assertion is used to check if an interface holds a specific type and
// then access its value
package main

import "fmt"

type Film struct {
	Name string
	Year int
}

func main() {
	var name any = "Jackson Lamb"

	{
		s, ok := name.(string)
		if ok {
			fmt.Printf("Slough House is run by %s\n", s)
		}
	}
	// Running a type assertion with 2 return values is the "safe" way of doing
	// this, because we avoid a panic when the type conversion fails.
	//
	// s := name.(int) would result in
	// "panic: interface conversion: interface {} is string, not int"
	{
		s, ok := name.(int)
		if !ok {
			fmt.Printf("%s is not an int and 's' is now %+v\n", name, s)
		}
	}
	{
		s, ok := name.(bool)
		if !ok {
			fmt.Printf("%s is not a bool and 's' is now %+v\n", name, s)
		}
	}
	{
		s, ok := name.(Film)
		if !ok {
			fmt.Printf("%s is not a Film and 's' is now %+v\n", name, s)
		}
	}

	assertType(name)
	assertType(42)
	assertType(42.0)
	assertType(Film{Name: "The Big Lebowski"})
}

func assertType(thing any) {
	// It's common to handle this scenario inside of a switch

	switch actual := thing.(type) {
	case string:
		fmt.Println("we have a string")
	case int:
		fmt.Println("we have an int")
	case float64:
		fmt.Println("we have a 64 bit float")
	default:
		fmt.Printf("we have something else: %+v\n", actual)
	}
}
