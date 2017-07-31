package atoms

import "encoding/json"

// Int is an atomic int
type Int struct {
	i Int64
}

// Load will get the current value
func (i *Int) Load() (n int) {
	return int(i.i.Load())
}

// Add will increment the current value by n
func (i *Int) Add(n int) (new int) {
	return int(i.i.Add(int64(n)))
}

// Store will perform an atomic store for a new value
func (i *Int) Store(new int) {
	i.i.Store(int64(new))
}

// Swap will perform an atomic swap for a new value
func (i *Int) Swap(new int) (old int) {
	return int(i.i.Swap(int64(new)))
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value
func (i *Int) CompareAndSwap(old, new int) (changed bool) {
	return i.i.CompareAndSwap(int64(old), int64(new))
}

// MarshalJSON is a json encoding helper function
func (i *Int) MarshalJSON() (b []byte, err error) {
	return json.Marshal(i.Load())
}

// UnmarshalJSON is a json decoding helper function
func (i *Int) UnmarshalJSON(b []byte) (err error) {
	var val int
	if err = json.Unmarshal(b, &val); err != nil {
		return
	}

	i.Store(val)
	return
}
