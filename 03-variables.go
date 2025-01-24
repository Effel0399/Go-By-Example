package main

import "fmt"

func main03() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	// variables with no initialization are zero-valued

	f := "apple"
	fmt.Println(f)
	// := short for declare and init
	// e.g. var f string = "apple"
	// only available in func
}
