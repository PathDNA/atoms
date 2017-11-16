package atoms

import (
	"sync"
	"testing"
)

func TestKeyLocker(t *testing.T) {
	var (
		val int
		wg  sync.WaitGroup
		mux KeyLocker
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mux.Update("key", func() {
			val++
			wg.Done()
		})
	}
	wg.Wait()

	if val != 10 {
		t.Fatalf(testErrInvalidValueFmt, 10, val)
	}
}

func TestRWKeyLocker(t *testing.T) {
	var (
		val  int
		sink int
		wg   sync.WaitGroup
		mux  KeyLocker
	)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go mux.Update("key", func() {
			val++
			wg.Done()
		})
		go mux.Read("key", func() {
			sink = val
			wg.Done()
		})
	}
	wg.Wait()

	if val != 10 {
		t.Fatalf(testErrInvalidValueFmt, 10, val)
	}

	_ = sink
}
