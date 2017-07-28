package atoms

import (
	"encoding/json"
	"testing"
)

func TestUint64(t *testing.T) {
	var u Uint64
	if val := u.Load(); val != 0 {
		t.Fatalf(testErrInvalidValueFmt, 0, val)
	}

	if val := u.Add(6); val != 6 {
		t.Fatalf(testErrInvalidValueFmt, 6, val)
	}

	if val := u.Add(7); val != 13 {
		t.Fatalf(testErrInvalidValueFmt, 13, val)
	}

	if val := u.Swap(3); val != 13 {
		t.Fatalf(testErrInvalidValueFmt, 13, val)
	}

	if u.CompareAndSwap(6, 13) {
		t.Fatalf(testErrInvalidSwapFmt)
	}
}

func TestUint64JSON(t *testing.T) {
	var ts testUint64
	b := []byte(`{ "number" : 7 }`)
	if err := json.Unmarshal(b, &ts); err != nil {
		t.Fatal(err)
	}

	if val := ts.Number.Load(); val != 7 {
		t.Fatalf(testErrInvalidValueFmt, 7, val)
	}
}

type testUint64 struct {
	Number Uint64 `json:"number"`
}
