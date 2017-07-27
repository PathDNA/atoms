package atoms

import "sync/atomic"

// Int32 is an atomic int32
type Int32 struct {
	v int32
}

// Get will get the current value
func (i *Int32) Get() (n int32) {
	return atomic.LoadInt32(&i.v)
}

// Add will add n to the current value
func (i *Int32) Add(n int32) (new int32) {
	return atomic.AddInt32(&i.v, n)
}

// Swap will perform an atomic swap for a new value
func (i *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(&i.v, new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (i *Int32) CompareAndSwap(old, new int32) (changed bool) {
	return atomic.CompareAndSwapInt32(&i.v, old, new)
}
