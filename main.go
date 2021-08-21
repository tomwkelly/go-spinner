package main

import (
	"fmt"
	"time"
)

var colours = map[string]string{
	"Reset": "0",
	"Black": "30",
	"Cyan":  "36",
}

func printIP(t string) {
	fmt.Printf("\r%s", t)
}

func hideCursor() {
	fmt.Printf("\x1b%s", "[?25l")
}

func showCursor() {
	fmt.Printf("\x1b%s", "[?25h")
}

func setColour(colour string) {
	printIP("\u001b[" + colours[colour] + "m")
}

func main() {
	hideCursor()
	frames := [10]string{"⠋",
		"⠙",
		"⠹",
		"⠸",
		"⠼",
		"⠴",
		"⠦",
		"⠧",
		"⠇",
		"⠏"}

	c := 0
	setColour("Cyan")
	for i := 0; i < 50; i++ {
		if c == len(frames) {
			c = 0
		}
		printIP(frames[c])
		time.Sleep(80 * time.Millisecond)
		c++
	}
	setColour("Reset")
	printIP("Done")
	showCursor()
	fmt.Println()
}
