package main

import "fmt"
import "sort"

var mp1 = make(map[int][]int, 100)

// LinkedHashMap is important!
type LinkedHashMap struct {
	key   string
	value int
}

func tryMap() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	_, isPresent := mapLit["three"]
	if !isPresent {
		mapLit["three"] = 5
	}
	_, isPresent = mapLit["three"]
	fmt.Println("key three is present?", isPresent)
	if _, ok := mapLit["three"]; ok {
		delete(mapLit, "three")
	}
	_, isPresent = mapLit["three"]
	fmt.Println("key three is present?", isPresent)

	for key, value := range mapLit {
		fmt.Println("key", key, "value", value)
	}

	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	sortMap()
}

func sortMap() {
	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	keys := make([]string, len(barVal))
	i := 0
	for key := range barVal {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
