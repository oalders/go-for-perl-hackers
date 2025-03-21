// package main employs several different sorting strategies.
package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	before := "abc123😊"
	fmt.Printf("%+v\n", viaSplit(before))
	fmt.Printf("%+v\n", viaRune(before))
	fmt.Printf("%+v\n", viaReverse(before))
	fmt.Printf("%+v\n", viaInPlace(before))
}

func viaSplit(before string) string {
	list := strings.Split(before, "")
	reverseList := []string{}
	for i := len(list) - 1; i >= 0; i-- {
		reverseList = append(reverseList, list[i])
	}
	return strings.Join(reverseList, "")
}

func viaRune(before string) string {
	reverseList := []rune{}
	for _, r := range before {
		reverseList = append([]rune{r}, reverseList...)
	}
	return string(reverseList)
}

// as of Go 1.21.
func viaReverse(before string) string {
	list := strings.Split(before, "")
	slices.Reverse(list)
	return strings.Join(list, "")
}

// lifted from slices.Reverse
// swaps positions:
// 0,6
// 1,5
// 2,4
func viaInPlace(before string) string {
	list := strings.Split(before, "")
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return strings.Join(list, "")
}
