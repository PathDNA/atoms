package atoms

import (
	"encoding/json"
	"testing"
)

func TestInt(t *testing.T) {
	var i Int
	if val := i.Load(); val != 0 {
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
}

func TestIntJSON(t *testing.T) {
	var ts testInt
	b := []byte(`{ "number" : 7 }`)
	if err := json.Unmarshal(b, &ts); err != nil {
		t.Fatal(err)
	}

	if val := ts.Number.Load(); val != 7 {
		t.Fatalf(testErrInvalidValueFmt, 7, val)
	}
}

type testInt struct {
	Number Int `json:"number"`
}
