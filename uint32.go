package atomic

import "sync/atomic"

// Uint32 is an atomic uint32
type Uint32 struct {
	v uint32
}

// Get will get the current value
func (u *Uint32) Get() (n uint32) {
	return atomic.LoadUint32(&u.v)
}

// Add will add n to the current value
func (u *Uint32) Add(n uint32) (new uint32) {
	return atomic.AddUint32(&u.v, n)
}

// Swap will perform an atomic swap for a new value
func (u *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(&u.v, new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (u *Uint32) CompareAndSwap(old, new uint32) (changed bool) {
	return atomic.CompareAndSwapUint32(&u.v, old, new)
}
