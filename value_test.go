package atoms

import (
	"encoding/json"
	"testing"
)

func TestValue(t *testing.T) {
	var v Value
	if val := v.Load(); val != nil {
		t.Fatalf(testErrInvalidValueFmt, nil, val)
	}

	v.Store(5)

	if val, ok := v.Swap(3).(int); !ok || val != 5 {
		t.Fatalf(testErrInvalidValueFmt, 5, val)
	}

	return
}

func TestValueJSON(t *testing.T) {
	var ts testValue
	b := []byte(`{ "v" : {"test": 5} }`)
	if err := json.Unmarshal(b, &ts); err != nil {
		t.Fatal(err)
	}
	if _, ok := ts.Value.Load().(map[string]interface{}); !ok {
		t.Fatalf(testErrInvalidTypeFmt, map[string]interface{}{}, ts.Value.Load())
	}

	ts.Value.Store(map[string]int{})
	if err := json.Unmarshal(b, &ts); err != nil {
		t.Fatal(err)
	}
	if m, ok := ts.Value.Load().(map[string]int); !ok || m["test"] != 5 {
		t.Errorf(testErrInvalidTypeFmt, map[string]int{}, ts.Value.Load())
		t.Fatalf(testErrInvalidValueFmt, map[string]int{"test": 5}, ts.Value.Load())
	}

}

type testValue struct {
	Value Value `json:"v"`
}
