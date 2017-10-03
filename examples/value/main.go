package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/PathDNA/atoms"
)

func main() {
	type dummy struct {
		V int `json:"v"`
	}
	var v atoms.Value

	// set the internal type
	v.Store(dummy{})

	b := []byte(`{ "v" : 45066 }`)
	if err := json.Unmarshal(b, &v); err != nil {
		log.Fatal(err)
	}
	v.CompareAndSwap(func(oldV interface{}) (newV interface{}, ok bool) {
		v, _ := oldV.(dummy)
		v.V++
		return v, true
	})

	dv, _ := v.Load().(dummy)

	fmt.Printf("%#+v\n", dv)
	fmt.Printf("0x%Xs\n", dv.V)
}
