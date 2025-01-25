package main

import (
	"fmt"
	"maps"
)

// Maps are Go’s built-in associative data type
// (sometimes called hashes or dicts in other languages).

func main10() {
	m := make(map[string]int) // create empty map "make(map[key-type]val-type)"

	m["k1"] = 7 // set key/value pairs with "name[key] = value"
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"] // get value for a key with "name[key]"
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3) // will return 0 if key not exist

	fmt.Println("len:", len(m)) // return the amount of key/value pair (not the length)

	delete(m, "k2") // delete key/value pairs
	fmt.Println("map:", m)

	clear(m) // remove ALL key/value from map
	fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2} // declare and init
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2} // maps syntax for equal
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
