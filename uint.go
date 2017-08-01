package atoms

import (
	"encoding/json"
	"sync/atomic"
)

// Uint is an atomic uint
type Uint struct {
	v uintptr
}

// Load will get the current value
func (u *Uint) Load() (n uint) {
	return uint(atomic.LoadUintptr(&u.v))
}

// Add will increment the current value by n
func (u *Uint) Add(n uint) (new uint) {
	return uint(atomic.AddUintptr(&u.v, uintptr(n)))
}

// Store will perform an atomic store for a new value
func (u *Uint) Store(new uint) {
	atomic.StoreUintptr(&u.v, uintptr(new))
}

// Swap will perform an atomic swap for a new value
func (u *Uint) Swap(new uint) (old uint) {
	return uint(atomic.SwapUintptr(&u.v, uintptr(new)))
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (u *Uint) CompareAndSwap(old, new uint) (changed bool) {
	return atomic.CompareAndSwapUintptr(&u.v, uintptr(old), uintptr(new))
}

// MarshalJSON is a json encoding helper function
func (u *Uint) MarshalJSON() (b []byte, err error) {
	return json.Marshal(u.Load())
}

// UnmarshalJSON is a json decoding helper function
func (u *Uint) UnmarshalJSON(b []byte) (err error) {
	var val uint
	if err = json.Unmarshal(b, &val); err != nil {
		return
	}

	u.Store(val)
	return
}
