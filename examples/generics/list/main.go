// Package main demonstrates how to use generics with slices and arrays. This
// is a more concise version of ./examples/data-structures/array/main.go.
package main

import (
	"fmt"
	"reflect"
)

// The type of a 5 element array is [5]int while the type of a slice is []int.
// The length and capacity of both slices and arrays can be inspected via the
// len() and cap() built-ins.
func main() {
	// even though list is created with 5 elements, it is a slice and not an array.
	list := make([]int, 5)
	inspect(list)
	list[3] = 100
	inspect(list)
	list = append(list, 88)
	inspect(list)

	// now create an array
	fixedList := [5]int{}

	fmt.Println("---")
	inspect(fixedList)

	// we cannot use append with an array, so we convert to a slice
	wasFixedList := fixedList[:]
	wasFixedList = append(wasFixedList, 88)
	inspect(wasFixedList)

	fmt.Println("---")

	// We can also convert to slice and append before variable assignment. This
	// is a little more succinct.
	oneStep := append(fixedList[:], 88)
	inspect(oneStep)

	inspect("foo")
}

// We don't save a lot of lines of code with generics here, but we now have one
// function rather than two. Also, we end up using reflection, which can be
// quite slow. For the purposes of this example reflection likely has little
// impact on the code's running time.
func inspect(list any) {
	v := reflect.ValueOf(list)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		fmt.Printf("len: %d\tcap: %d\n", v.Len(), v.Cap())
	} else {
			fmt.Printf("💥 %+v is of the type %s\n", list, v.Kind())
	}
}
