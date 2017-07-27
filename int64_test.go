package atoms

import (
	"testing"
)

func TestInt64(t *testing.T) {
	var i Int64
	if val := i.Get(); val != 0 {
		t.Fatalf(testErrInvalidValueFmt, 0, val)
	}

	if val := i.Add(6); val != 6 {
		t.Fatalf(testErrInvalidValueFmt, 6, val)
	}

	if val := i.Add(7); val != 13 {
		t.Fatalf(testErrInvalidValueFmt, 13, val)
	}

	if val := i.Swap(3); val != 13 {
		t.Fatalf(testErrInvalidValueFmt, 13, val)
	}

	if i.CompareAndSwap(6, 13) {
		t.Fatalf(testErrInvalidSwapFmt)
	}

	return
}
