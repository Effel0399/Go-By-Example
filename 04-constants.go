package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main04() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	// numeric constant has no type until given one
	// such as by an explicit conversion

	fmt.Println(math.Sin(n))
	// a number can be given a type in context that requires it
	// example, math.Sin expects a float64
}
