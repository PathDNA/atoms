package atomic

import (
	"testing"
)

func TestUint32(t *testing.T) {
	var u Uint32
	if val := u.Get(); val != 0 {
		t.Fatalf("Invalid value, expected %d and received %d", 0, val)
	}

	if val := u.Add(6); val != 6 {
		t.Fatalf("Invalid value, expected %d and received %d", 6, val)
	}

	if val := u.Add(7); val != 13 {
		t.Fatalf("Invalid value, expected %d and received %d", 13, val)
	}

	if val := u.Swap(3); val != 13 {
		t.Fatalf("Invalid value, expected %d and received %d", 13, val)
	}

	if u.CompareAndSwap(6, 13) {
		t.Fatalf("Swapped successfully when should have failed")
	}

	return
}
