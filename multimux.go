package atoms

import "sync"

// MultiMux allows to hold locks by a key.
type MultiMux struct {
	mux sync.RWMutex
	m   map[string]*sync.RWMutex
}

// Update executes fn while the mutex is write-locked and guarantees the mutex is released even in the case of a panic.
func (mm *MultiMux) Update(key string, fn func()) {
	m := mm.RWMutex(key)
	m.Lock()
	defer m.Unlock()

	fn()
}

// Read executes fn while the mutex is read-locked and guarantees the mutex is released even in the case of a panic.
func (mm *MultiMux) Read(key string, fn func()) {
	m := mm.RWMutex(key)
	m.RLock()
	defer m.RUnlock()

	fn()
}

// RWMutex returns a RWMutex for the given key.
func (mm *MultiMux) RWMutex(key string) (m *sync.RWMutex) {
	if m = mm.get(key); m != nil {
		return
	}

	mm.mux.Lock()
	if mm.m == nil {
		mm.m = map[string]*sync.RWMutex{}
	}
	if m = mm.m[key]; m == nil {
		m = new(sync.RWMutex)
		mm.m[key] = m
	}
	mm.mux.Unlock()

	return
}

// Delete deletes the specific key, ensuring that it isn't used anywhere else.
func (mm *MultiMux) Delete(key string) {
	m := mm.get(key)
	if m == nil { // lock doesn't exist, we can just return
		return
	}

	m.Lock() // acquire the lock first to make sure it's not used somewhere else
	mm.mux.Lock()
	delete(mm.m, key)
	mm.mux.Unlock()
	m.Unlock()
}

func (mm *MultiMux) get(key string) *sync.RWMutex {
	mm.mux.RLock()
	m := mm.m[key]
	mm.mux.RUnlock()
	return m
}
