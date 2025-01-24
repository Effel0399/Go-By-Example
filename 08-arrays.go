package main

import "fmt"

func main08() {
	var a [5]int // the array will hold 5 ints
	fmt.Println("emp:", a)

	a[4] = 100 // set the 4th array to value
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) // print the 4th array

	fmt.Println("len:", len(a)) // get the length of array

	b := [5]int{1, 2, 3, 4, 5} // declare and init array
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} // let the compiler count the element
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} // specify index with :, the element between will be zeroed
	fmt.Println("idx:", b)

	var twoD [2][3]int // multi-dimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // init multi-dimensional array at once

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
