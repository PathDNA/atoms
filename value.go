package atoms

import (
	"encoding/json"
	"reflect"
	"sync"
)

// Value is an atomic interface{}
type Value struct {
	m sync.RWMutex
	v interface{}
}

// Load will get the current value
func (av *Value) Load() (v interface{}) {
	av.m.RLock()
	v = av.v
	av.m.RUnlock()
	return
}

// Store will perform an atomic store for a new value
func (av *Value) Store(v interface{}) {
	av.m.Lock()
	av.v = v
	av.m.Unlock()
}

// Swap will perform an atomic swap for a new value and return the old value
func (av *Value) Swap(newV interface{}) (oldV interface{}) {
	av.m.Lock()
	oldV, av.v = av.v, newV
	av.m.Unlock()
	return
}

// CompareAndSwap will perform an atomic compare and swap for an old and new value,
// fn is expected to check the old value and return the new value and true, or return false
func (av *Value) CompareAndSwap(fn func(oldV interface{}) (newV interface{}, ok bool)) (ok bool) {
	var newV interface{}
	av.m.Lock()
	if newV, ok = fn(av.v); ok {
		av.v = newV
	}
	av.m.Unlock()
	return
}

// MarshalJSON is a json encoding helper function
func (av *Value) MarshalJSON() (b []byte, err error) {
	return json.Marshal(av.Load())
}

// UnmarshalJSON is a json decoding helper function,
// to specify the type of the value make sure to call .Set with the zero type you want.
// Example:
// 	v.Set(&Struct{})
// err := json.Unmarshal(data, &v)
func (av *Value) UnmarshalJSON(b []byte) (err error) {
	if oval := av.Load(); oval != nil {
		var (
			typ = reflect.TypeOf(oval)
			val = reflect.New(typ)
		)
		if err = json.Unmarshal(b, val.Interface()); err != nil {
			return
		}
		av.Store(val.Elem().Interface())
	} else {
		var val interface{}
		if err = json.Unmarshal(b, &val); err != nil {
			return
		}
		av.Store(val)
	}
	return
}
