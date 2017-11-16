package atoms

import "sync"

// KeyLocker allows to hold locks by key.
type KeyLocker struct {
	m   map[string]*sync.RWMutex
	mux sync.RWMutex
}

// RWMutex returns a RWMutex for the specifc key.
func (kl *KeyLocker) RWMutex(key string) (m *sync.RWMutex) {
	kl.mux.RLock()
	m = kl.m[key]
	kl.mux.RUnlock()
	if m != nil {
		return
	}

	kl.mux.Lock()
	if kl.m == nil {
		kl.m = map[string]*sync.RWMutex{}
	}
	if m = kl.m[key]; m == nil {
		m = new(sync.RWMutex)
		kl.m[key] = m
	}
	kl.mux.Unlock()

	return
}

func (kl *KeyLocker) Delete(key string) {
	kl.mux.RLock()
	m := kl.m[key]
	kl.mux.RUnlock()
	if m == nil { // lock doesn't exist, we can just return
		return
	}

	m.Lock() // acquire the lock first to make sure it's not used somewhere else
	kl.mux.Lock()
	delete(kl.m, key)
	kl.mux.Unlock()
	m.Unlock()
}

// Lock Locks the specific key and returns the Unlock function.
// Example: defer l.Lock("key")()
func (kl *KeyLocker) Lock(key string) func() {
	m := kl.RWMutex(key)
	m.Lock()
	return m.Unlock
}

// RLock RLocks the specific key and returns the RUnlock function.
// Example: defer l.RLock("key")()
func (kl *KeyLocker) RLock(key string) func() {
	m := kl.RWMutex(key)
	m.RLock()
	return m.RUnlock
}

// Update executes fn while the mutex is write-locked and guarantees the mutex is released even in the case of a panic.
func (kl *KeyLocker) Update(key string, fn func()) {
	defer kl.Lock(key)()

	fn()
}

// Read executes fn while the mutex is read-locked and guarantees the mutex is released even in the case of a panic.
func (kl *KeyLocker) Read(key string, fn func()) {
	defer kl.RLock(key)()

	fn()
}
