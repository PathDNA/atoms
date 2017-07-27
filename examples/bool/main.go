package main

import (
	"fmt"
	"github.com/Path94/atoms"
)

func main() {
	var b atoms.Bool
	// Get current state
	state := b.Get()
	fmt.Printf("State: %v\n", state)

	// Set value, will fail because value has not changed
	changed := b.Set(false)
	fmt.Printf("Changed: %v\n", changed)

	// Set value to true
	changed = b.Set(true)
	fmt.Printf("Changed: %v\n", changed)

	// Get current state
	state = b.Get()
	fmt.Printf("State: %v\n", state)
}
