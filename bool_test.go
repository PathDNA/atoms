package atoms

import (
	"encoding/json"
	"testing"
)

func TestBool(t *testing.T) {
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

func TestBoolJSON(t *testing.T) {
	var ts testBool
	b := []byte(`{ "state" : true }`)
	if err := json.Unmarshal(b, &ts); err != nil {
		t.Fatal(err)
	}

	if !ts.State.Get() {
		t.Fatal("received a false negative")
	}
}

type testBool struct {
	State Bool `json:"state"`
}
