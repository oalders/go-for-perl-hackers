// Package main demonstrates the use of "errors.Is".
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	filepath := "arrested-development.mov"

	file, err := openFile(filepath)

	// Demonstrate how to use errors.Is to check for a sentinel error.
	if err != nil {
		fmt.Printf("Full error: %s\n", err)

		var targetErr *os.PathError

		// errors.As() searches from the outermost error (the one you provide
		// directly) inward through the chain. It looks at each error in
		// sequence as it unwraps them:
		//
		// First, it checks if the outermost error matches the target type. If
		// not, it unwraps the error and checks that one. It continues this
		// process until it either finds a match or reaches the end of the
		// chain.
		//
		// So "first matching error" means the first one encountered when
		// moving from the outermost (most recently wrapped) error toward the
		// innermost (original base) error.
		if as := errors.As(err, &targetErr); as {
			fmt.Printf("Base error: %s [via errors.As]\n", targetErr)
		}

		// errors.Is searches in a similar fashion to errors.As It begins with
		// the outermost error and continues to unwrap until it finds a match
		// or reaches the innermost error.
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("This is an os.ErrNotExist sentinel error")
		}

		// Get the the innermost error by unwrapping.
		inner := errors.Unwrap(err)
		fmt.Printf("Base error: %s [via errors.Unwrap]\n", inner)
	} else {
		fmt.Printf("opened file: %s\n", file.Name())
	}
}

// Use fmt.Errorf to add some more context to the original error.
func openFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath) // For read access.
	if err != nil {
		return nil, fmt.Errorf("💥: %w", err)
	}
	return file, nil
}
