package atoms

import "sync"

// Mux wraps a sync.Mutex to allow simple and safe operation on the mutex.
type Mux struct {
	mux sync.Mutex
}

// Do executes fn while the mutex is locked and guarantees the mutex is released even in the case of a panic.
func (m *Mux) Do(fn func()) {
	m.mux.Lock()
	defer m.mux.Unlock()
	fn()
}

// RWMux wraps a sync.RWMutex to allow simple and safe operation on the mutex.
type RWMux struct {
	mux sync.RWMutex
}

// Update executes fn while the mutex is write-locked and guarantees the mutex is released even in the case of a panic.
func (m *RWMux) Update(fn func()) {
	m.mux.Lock()
	defer m.mux.Unlock()
	fn()
}

// Read executes fn while the mutex is read-locked and guarantees the mutex is released even in the case of a panic.
func (m *RWMux) Read(fn func()) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	fn()
}
