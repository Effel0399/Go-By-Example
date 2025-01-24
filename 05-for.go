package main

import "fmt"

func main05() {
	// basic loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// init>condition>after loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// loop using range over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}

	// loop infinite until break
	for {
		fmt.Println("loop")
		break
	}

	// use continue to the next iteration of the loop
	for n := range 6 {
		if n%2 == 1 {
			continue
		}
		fmt.Println(n)
	}
}
