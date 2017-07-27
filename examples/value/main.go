package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Path94/atoms"
)

func main() {
	type dummy struct {
		V int `json:"v"`
	}
	var v atoms.Value

	// set the internal type
	v.Store(&dummy{})

	b := []byte(`{ "v" : 45067 }`)
	if err := json.Unmarshal(b, &v); err != nil {
		log.Fatal(err)
	}
	dv, _ := v.Load().(*dummy)

	fmt.Printf("%#+v\n", dv)
	fmt.Printf("0x%Xs\n", dv.V)
}
