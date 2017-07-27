package atoms

import (
	"encoding/json"
	"sync/atomic"
)

// Bool is an atomic bool type
type Bool struct {
	v int32
}

const (
	// False state
	False int32 = iota
	// True state
	True
)

// Get will get the current state
func (b *Bool) Get() (state bool) {
	return atomic.LoadInt32(&b.v) == True
}

// Set will set the state
func (b *Bool) Set(state bool) (changed bool) {
	if state {
		return atomic.CompareAndSwapInt32(&b.v, False, True)
	}

	return atomic.CompareAndSwapInt32(&b.v, True, False)
}

// MarshalJSON is a json encoding helper function
func (b *Bool) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(b.Get())
}

// UnmarshalJSON is a json decoding helper function
func (b *Bool) UnmarshalJSON(bs []byte) (err error) {
	var val bool
	if err = json.Unmarshal(bs, &val); err != nil {
		return
	}

	b.Set(val)
	return
}
