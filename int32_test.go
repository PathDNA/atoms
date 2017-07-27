package atomic

import (
	"testing"
)

func TestInt32(t *testing.T) {
	var i Int32
	if val := i.Get(); val != 0 {
		t.Fatalf("Invalid value, expected %d and received %d", 0, val)
	}

	if val := i.Add(6); val != 6 {
		t.Fatalf("Invalid value, expected %d and received %d", 6, val)
	}

	if val := i.Add(7); val != 13 {
		t.Fatalf("Invalid value, expected %d and received %d", 13, val)
	}

	if val := i.Swap(3); val != 13 {
		t.Fatalf("Invalid value, expected %d and received %d", 13, val)
	}

	if i.CompareAndSwap(6, 13) {
		t.Fatalf("Swapped successfully when should have failed")
	}

	return
}
