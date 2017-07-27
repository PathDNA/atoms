# Atoms
Atoms is a QoL helper library which provides atomic primitives. 

## Purpose
The goal of this library is to reduce the amount of mundane boilplate code associated with managing atomic-friendly values

## Usage
### Int64
```go
package main

import (
	"fmt"
	"github.com/Path94/atoms"
)

func main() {
	var i atoms.Int64
	// Set value to 7
	i.Store(7)
	current := i.Load()
	fmt.Printf("Current value: %d\n", current)

	// Swap value with 13, returned value will be our old value
	old := i.Swap(13)
	fmt.Printf("Old value: %d\n", old)

	// Increment value by 6, returned value will be our new value
	new := i.Add(6)
	fmt.Printf("New value: %d\n", new)

	// Compare and swap value, will fail because value is 19 (not 20)
	changed := i.CompareAndSwap(20, 40)
	fmt.Printf("Changed: %v\n", changed)
}

```

### Bool
```go
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

```

*Note - Check out the examples directory for compilable examples*