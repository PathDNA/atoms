# Atoms [![GoDoc](https://godoc.org/github.com/Path94/atoms?status.svg)](https://godoc.org/github.com/Path94/atoms) ![Status](https://img.shields.io/badge/status-beta-yellow.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/Path94/atoms)](https://goreportcard.com/report/github.com/Path94/atoms)

Atoms is a QoL helper library which provides atomic primitives. The goal of this library is to reduce the amount of mundane boilplate code associated with managing atomic-friendly values

### Provided primitives
- int32
- int64
- uint32
- uint64
- boolean
- Value (interface{})
## Features
### Numeric values
- Load
- Store
- Swap
- CompareAndSwap
- JSON Marshal/Unmarshal

### Boolean
- Get
- Set (Compare and swap functionality)
- JSON Marshal/Unmarshal

### Generic
- Get
- Set (Compare and swap functionality)
- JSON Marshal/Unmarshal with optional types.

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

### Value
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Path94/atoms"
)

func main() {
	type dummy struct {
		V int `json:"v"`
	}
	var v atoms.Value

	// set the internal type
	v.Store(dummy{})

	b := []byte(`{ "v" : 45066 }`)
	if err := json.Unmarshal(b, &v); err != nil {
		log.Fatal(err)
	}
	v.CompareAndSwap(func(oldV interface{}) (newV interface{}, ok bool) {
		v, _ := oldV.(dummy)
		v.V++
		return v, true
	})

	dv, _ := v.Load().(dummy)

	fmt.Printf("%#+v\n", dv)
	fmt.Printf("0x%Xs\n", dv.V)
}
```

*Note - Check out the examples directory for compilable examples*
