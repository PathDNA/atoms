package atoms

import (
	"encoding/json"
	"sync/atomic"
)

// Uint32 is an atomic uint32
type Uint32 struct {
	v uint32
}

// Load will get the current value
func (u *Uint32) Load() (n uint32) {
	return atomic.LoadUint32(&u.v)
}

// Add will increment the current value by n
func (u *Uint32) Add(n uint32) (new uint32) {
	return atomic.AddUint32(&u.v, n)
}

// Store will perform an atomic store for a new value
func (u *Uint32) Store(new uint32) {
	atomic.StoreUint32(&u.v, new)
}

// Swap will perform an atomic swap for a new value
func (u *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(&u.v, new)
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (u *Uint32) CompareAndSwap(old, new uint32) (changed bool) {
	return atomic.CompareAndSwapUint32(&u.v, old, new)
}

// MarshalJSON is a json encoding helper function
func (u *Uint32) MarshalJSON() (b []byte, err error) {
	return json.Marshal(u.Load())
}

// UnmarshalJSON is a json decoding helper function
func (u *Uint32) UnmarshalJSON(b []byte) (err error) {
	var val uint32
	if err = json.Unmarshal(b, &val); err != nil {
		return
	}

	u.Store(val)
	return
}
