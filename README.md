# Atoms [![GoDoc](https://godoc.org/github.com/Path94/atoms?status.svg)](https://godoc.org/github.com/Path94/atoms) ![Status](https://img.shields.io/badge/status-beta-yellow.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/Path94/atoms)](https://goreportcard.com/report/github.com/Path94/atoms)

Atoms is a QoL helper library which provides atomic primitives. The goal of this library is to reduce the amount of mundane boilplate code associated with managing atomic-friendly values

### Provided primitives
- int
- uint
- int32
- int64
- uint32
- uint64
- boolean
- string

### Non-primitive helpers
- generic value (interface{})
- mutex wrapper
- rwmutex wrapper

## Features
### Numeric values
- Load
- Store
- Swap
- Add
- CompareAndSwap
- JSON Marshal/Unmarshal

### Boolean
- Get
- Set (Compare and swap functionality)
- JSON Marshal/Unmarshal

### String
- Load
- Store
- Swap
- JSON

### Generic
- Load
- Store
- Swap
- CompareAndSwap
- JSON Marshal/Unmarshal

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

### String
```go
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

### Mux/RWMux
```go
package main

import (
	"fmt"
	"sync"

	"github.com/Path94/atoms"
)

func main() {
	var wg sync.WaitGroup
	c := NewCounter()
	wg.Add(3)

	go func() {
		c.Increment("foo")
		wg.Done()
	}()

	go func() {
		c.Increment("foo")
		wg.Done()
	}()

	go func() {
		c.Increment("foo")
		wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Output will be 3
	fmt.Println(c.Get("foo"))
}

// NewCounter will return a new counter
func NewCounter() *Counter {
	var c Counter
	// Initialize our internal map
	c.cm = make(map[string]uint64)
	return &c
}

// Counter manages a counter
type Counter struct {
	mux atoms.RWMux
	cm  map[string]uint64
}

// Increment will increment a given key
func (c *Counter) Increment(key string) {
	c.mux.Update(func() {
		c.cm[key]++
	})
}

// Get will get the counter value for a given key
func (c *Counter) Get(key string) (n uint64) {
	c.mux.Read(func() {
		n = c.cm[key]
	})

	return
}

```
*Note - Check out the examples directory for compilable examples*
