package atoms

import "sync/atomic"

// Int64 is an atomic int64
type Int64 struct {
	v int64
}

// Load will get the current value
func (i *Int64) Load() (n int64) {
	return atomic.LoadInt64(&i.v)
}

// Add will increment the current value by n
func (i *Int64) Add(n int64) (new int64) {
	return atomic.AddInt64(&i.v, n)
}

// Swap will perform an atomic swap for a new value
func (i *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64(&i.v, new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (i *Int64) CompareAndSwap(old, new int64) (changed bool) {
	return atomic.CompareAndSwapInt64(&i.v, old, new)
}
