package atoms

import "testing"

func TestAtomicBool(t *testing.T) {
	var b Bool
	if b.Get() {
		t.Fatal("invalid state")
	}

	if !b.Set(true) {
		t.Fatal("state change not triggered")
	}

	if !b.Get() {
		t.Fatal("invalid state")
	}

	if b.Set(true) {
		t.Fatal("state change triggered when it shouldn't have")
	}

	if !b.Set(false) {
		t.Fatal("state change not triggered")
	}

	if b.Get() {
		t.Fatal("invalid state")
	}
}
