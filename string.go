package atoms

import (
	"encoding/json"
	"sync/atomic"
	"unsafe"
)

// String is a string type
type String struct {
	p unsafe.Pointer
}

// Load will get the current value
func (s *String) Load() (str string) {
	if up := atomic.LoadPointer(&s.p); up != nil {
		str = *(*string)(up)
	}

	return
}

// Store will perform an atomic store for a new value
func (s *String) Store(new string) {
	atomic.StorePointer(&s.p, unsafe.Pointer(&new))
}

// Swap will perform an atomic swap for a new value
func (s *String) Swap(new string) (old string) {
	if up := atomic.SwapPointer(&s.p, unsafe.Pointer(&new)); up != nil {
		old = *(*string)(up)
	}

	return
}

// MarshalJSON is a json encoding helper function
func (s *String) MarshalJSON() (b []byte, err error) {
	return json.Marshal(s.Load())
}

// UnmarshalJSON is a json decoding helper function
func (s *String) UnmarshalJSON(b []byte) (err error) {
	var val string
	if err = json.Unmarshal(b, &val); err != nil {
		return
	}

	s.Store(val)
	return
}
