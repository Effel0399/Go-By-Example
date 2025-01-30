package main

import "fmt"

// intSeq returns another fucntion
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main14() {

	// call intSeq, assign result to nextInt
	// this func captures its own i value
	// which will be updated each time we call nextInt
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
