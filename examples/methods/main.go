// package main demonstrates 3 different ways in which a method can be called.
// * Method Call
// * Method Value
// * Method Expression
package main

import "fmt"

type Drummer struct {
	Emoji      string
	Instrument string
	Job        string
}

func (d Drummer) Perform() {
	fmt.Printf("%s The %s plays the %s\n", d.Emoji, d.Job, d.Instrument)
}

func main() {
	neil := Drummer{Emoji: "🥁", Instrument: "drums", Job: "drummer"}

	// Method Call
	neil.Perform()

	// We can store the func in a variable before calling it.
	p := neil.Perform

	// Method Value
	p()

   perform := Drummer.Perform

   // Method Expression
   perform(neil)

   // Method Expression in a single line.
   Drummer.Perform(neil)
}
