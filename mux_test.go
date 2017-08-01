package atoms

import (
	"sync"
	"testing"
)

func TestMux(t *testing.T) {
	var (
		val int
		wg  sync.WaitGroup
		mux Mux
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mux.Do(func() {
			val++
			wg.Done()
		})
	}
	wg.Wait()

	if val != 10 {
		t.Fatalf(testErrInvalidValueFmt, 10, val)
	}
}
