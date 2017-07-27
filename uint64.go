package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Uint64 is an atomic uint64
type Uint64 uint64

func (u *Uint64) getIntPtr() *uint64 {
	return (*uint64)(unsafe.Pointer(u))
}

// Get will get the current value
func (u *Uint64) Get() (n uint64) {
	return atomic.LoadUint64(u.getIntPtr())
}

// Add will add n to the current value
func (u *Uint64) Add(n uint64) (new uint64) {
	return atomic.AddUint64(u.getIntPtr(), n)
}

// Swap will perform an atomic swap for a new value
func (u *Uint64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64(u.getIntPtr(), new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (u *Uint64) CompareAndSwap(old, new uint64) (changed bool) {
	return atomic.CompareAndSwapUint64(u.getIntPtr(), old, new)
}
