package atoms

import (
	"encoding/json"
	"testing"
)

func TestString(t *testing.T) {
	var s String
	s.Store("Hello world")
	if str := s.Load(); str != "Hello world" {
		t.Fatalf("invalid string, expected \"%s\" and received \"%s\"", "Hello world", str)
	}

	if str := s.Swap("Goodbye world"); str != "Hello world" {
		t.Fatalf("invalid string, expected \"%s\" and received \"%s\"", "Hello world", str)
	}

	if str := s.Load(); str != "Goodbye world" {
		t.Fatalf("invalid string, expected \"%s\" and received \"%s\"", "Goodbye world", str)
	}
}

func TestStringJSON(t *testing.T) {
	var ts testString
	b := []byte(`{ "name" : "John Doe" }`)
	if err := json.Unmarshal(b, &ts); err != nil {
		t.Fatal(err)
	}

	if val := ts.Name.Load(); val != "John Doe" {
		t.Fatalf(testErrInvalidValueFmt, "John Doe", val)
	}
}

type testString struct {
	Name String `json:"name"`
}
