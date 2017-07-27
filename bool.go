package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Bool is an atomic bool type
type Bool int32

const (
	// False state
	False int32 = iota
	// True state
	True
)

func (b *Bool) getIntPtr() *int32 {
	return (*int32)(unsafe.Pointer(b))
}

// Get will get the current state
func (b *Bool) Get() (state bool) {
	return atomic.LoadInt32(b.getIntPtr()) == True
}

// Set will set the state
func (b *Bool) Set(state bool) (changed bool) {
	if state {
		return atomic.CompareAndSwapInt32(b.getIntPtr(), False, True)
	}

	return atomic.CompareAndSwapInt32(b.getIntPtr(), True, False)
}
