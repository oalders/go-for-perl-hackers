package main

import "fmt"

// The only requirement of the Musician interface is that a Perform() function
// must be implemented.
type Musician interface {
	Perform()
}

type Drummer struct {
	Emoji      string
	Instrument string
	Job        string
}

type Singer struct {
	Emoji      string
	Instrument string
	Job        string
}

func (d Drummer) Perform() {
	fmt.Printf("%s The %s plays the %s\n", d.Emoji, d.Job, d.Instrument)
}

// play expects something that implements the Musician interface
func play(m Musician) {
	m.Perform()
}

func main() {
	neil := Drummer{Emoji: "🥁", Instrument: "drums", Job: "drummer"}
	play(neil)

	// If we uncomment the following code, it will throw the error:

	//"cannot use geddy (variable of struct type Singer) as Musician value in
	//argument to play: Singer does not implement Musician (missing method
	//Perform)"

	// geddy := Singer{Emoji: "🎙️", Instrument: "voice", Job: "singer"}
	// play(geddy)
}
