package atomic

import "sync/atomic"

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
