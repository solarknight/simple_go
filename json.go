package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func tryJson() {
	toJson()
}

func toJson() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)
}

func fromJson() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	// unmarshal
	js2, _ := json.Marshal(pa)
	fmt.Printf("Pointer JSON format: %s\n", js2)
	var data Address
	json.Unmarshal(js2, &data)
	fmt.Println("Revert data", data)
}
