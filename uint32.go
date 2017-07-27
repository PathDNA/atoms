package atomic

import (
	"sync/atomic"
	"unsafe"
)

// Uint32 is an atomic uint32
type Uint32 uint32

func (u *Uint32) getIntPtr() *uint32 {
	return (*uint32)(unsafe.Pointer(u))
}

// Get will get the current value
func (u *Uint32) Get() (n uint32) {
	return atomic.LoadUint32(u.getIntPtr())
}

// Add will add n to the current value
func (u *Uint32) Add(n uint32) (new uint32) {
	return atomic.AddUint32(u.getIntPtr(), n)
}

// Swap will perform an atomic swap for a new value
func (u *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(u.getIntPtr(), new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (u *Uint32) CompareAndSwap(old, new uint32) (changed bool) {
	return atomic.CompareAndSwapUint32(u.getIntPtr(), old, new)
}
