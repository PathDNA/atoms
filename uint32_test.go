package atoms

import (
	"encoding/json"
	"testing"
)

var testUint32Value uint32

func TestUint32(t *testing.T) {
	var u Uint32
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

func BenchmarkUint32(b *testing.B) {
	var u Uint32
	for i := 0; i < b.N; i++ {
		u.Add(1)
	}

	testUint32Value = u.Load()
}

func TestUint32JSON(t *testing.T) {
	var ts testUint32
	b := []byte(`{ "number" : 7 }`)
	if err := json.Unmarshal(b, &ts); err != nil {
		t.Fatal(err)
	}

	if val := ts.Number.Load(); val != 7 {
		t.Fatalf(testErrInvalidValueFmt, 7, val)
	}
}

type testUint32 struct {
	Number Uint32 `json:"number"`
}
