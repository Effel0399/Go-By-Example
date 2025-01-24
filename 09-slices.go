package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string // unitialized slice has 0 length and equal to nil
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // create empty slice with make
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a" // set and get slices
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // return slice length

	s = append(s, "d") // add an element to slices
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // make and copy slices
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] // slice[low:high] will show element s[2], s[3], and s[4]
	fmt.Println("sl1:", l)

	l = s[:5] // only slice up to and excluding s[5]
	fmt.Println("sl2:", l)

	l = s[2:] // only slice up from and including s[2]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // declare and init slices
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"} // print if two sets of slices has equal element
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) // multi-dimensional slices can vary in length
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
