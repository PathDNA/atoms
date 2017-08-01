package main

import (
	"fmt"
	"github.com/Path94/atoms"
)

func main() {
	var s atoms.String
	// Set value to "Hello world"
	s.Store("Hello world")
	current := s.Load()
	fmt.Printf("Current value: %s\n", current)

	// Swap value with "Goodbye world", returned value will be our old value
	old := s.Swap("Goodbye world")
	fmt.Printf("Old value: %s\n", old)

	current := s.Load()
	fmt.Printf("New current value: %s\n", current)
}
