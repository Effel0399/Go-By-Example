package main

import "fmt"

// This feature is used often in idiomatic Go
// for example to return both result and error values from a function.

func vals() (int, int) { // the function returns 2 ints.
	return 3, 70
}

func main12() {
	a, b := vals()
	fmt.Println(a) // return 1 value
	fmt.Println(b)

	fmt.Println(vals()) // return all

	_, c := vals() // return only subset
	fmt.Println(c)
}
