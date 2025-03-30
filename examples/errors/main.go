// Package main demonstrates the use of "errors.Is".
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	filepath := "somefile.go"
	file, err := os.Open(filepath) // For read access.
	if err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Printf("file %s does not exist\n", filepath)
	} else {
		fmt.Printf("could open %s for reading", file.Name())
	}
}
