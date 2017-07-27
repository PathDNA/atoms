package atoms

import (
	"encoding/json"
	"sync/atomic"
)

// Int32 is an atomic int32
type Int32 struct {
	v int32
}

// Load will get the current value
func (i *Int32) Load() (n int32) {
	return atomic.LoadInt32(&i.v)
}

// Add will increment the current value by n
func (i *Int32) Add(n int32) (new int32) {
	return atomic.AddInt32(&i.v, n)
}

// Store will perform an atomic store for a new value
func (i *Int32) Store(new int32) {
	atomic.StoreInt32(&i.v, new)
}

// Swap will perform an atomic swap for a new value
func (i *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(&i.v, new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (i *Int32) CompareAndSwap(old, new int32) (changed bool) {
	return atomic.CompareAndSwapInt32(&i.v, old, new)
}

// MarshalJSON is a json encoding helper function
func (i *Int32) MarshalJSON() (b []byte, err error) {
	return json.Marshal(i.Load())
}

// UnmarshalJSON is a json decoding helper function
func (i *Int32) UnmarshalJSON(b []byte) (err error) {
	var val int32
	if err = json.Unmarshal(b, &val); err != nil {
		return
	}

	i.Store(val)
	return
}
