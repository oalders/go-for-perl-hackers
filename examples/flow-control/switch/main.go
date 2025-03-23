// package main demonstrates a switch statements with and without expressions.
// emulate different editors via:
// EDITOR=nano go run main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	editor()

	currentCharge := []int{95, 65, 45, 25, 5, 3}
	for _, v := range currentCharge {
		battery(v)
	}
}

func editor() {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		fmt.Println("EDITOR env not set")
		os.Exit(1)
	}

	// default: does not need to come last
	switch editor {
	default:
		fmt.Printf(
			"I don't know where the source code for \"%s\" is\n",
			editor,
		)
	case "nvim":
		fmt.Println("https://github.com/neovim/neovim")
	case "vim":
		fmt.Println("https://github.com/vim/vim")
	case "nano":
		fmt.Println("https://www.nano-editor.org/git.php")
	}
}

func battery(charge int) {
	state := ""

	// use switch without an expression
	switch {
	case charge > 90:
		state = "is mostly fully charged"
	case charge > 60:
		state = "is ok"
	case charge > 30:
		state = "needs charge soon"
	case charge >= 5:
		state = "close to shutdown"
	default:
		state = "essentially empty"
	}
	fmt.Printf(
		"current battery charge %d means battery %s\n",
		charge,
		state,
	)
}
