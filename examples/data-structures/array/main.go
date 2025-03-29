package main

import (
	"fmt"
)

// The type of a 5 element array is [5]int while the type of a slice is []int.
// The length and capacity of both slices and arrays can be inspected via the
// len() and cap() built-ins.
func main() {
	// even though list is created with 5 elements, it is a slice and not an array.
	list := make([]int, 5)
	inspectSlice(list)
	list[3] = 100
	inspectSlice(list)
	list = append(list, 88)
	inspectSlice(list)

	// now create an array
	fixedList := [5]int{}

	fmt.Println("---")
	inspectArray(fixedList)

	// we cannot use append with an array, so we convert to a slice
	wasFixedList := fixedList[:]
	wasFixedList = append(wasFixedList, 88)
	inspectSlice(wasFixedList)

	fmt.Println("---")

	// We can also convert to slice and append before variable assignment. This
	// is a little more succinct.
	oneStep := append(fixedList[:], 88)
	inspectSlice(oneStep)
}

// Having two functions that do essentially the same thing is a good case for
// generics. See ./examples/generics/list
func inspectSlice (list []int ) {
	fmt.Printf("len: %d\tcap: %d\n", len(list), cap(list))
}

func inspectArray (list [5]int ) {
	fmt.Printf("len: %d\tcap: %d\n", len(list), cap(list))
}
