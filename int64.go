package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Int64 is an atomic int64
type Int64 int64

func (i *Int64) getIntPtr() *int64 {
	return (*int64)(unsafe.Pointer(i))
}

// Get will get the current value
func (i *Int64) Get() (n int64) {
	return atomic.LoadInt64(i.getIntPtr())
}

// Add will add n to the current value
func (i *Int64) Add(n int64) (new int64) {
	return atomic.AddInt64(i.getIntPtr(), n)
}

// Swap will perform an atomic swap for a new value
func (i *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64(i.getIntPtr(), new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (i *Int64) CompareAndSwap(old, new int64) (changed bool) {
	return atomic.CompareAndSwapInt64(i.getIntPtr(), old, new)
}
