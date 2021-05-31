package main

import "fmt"

// A variadic parameter is one that takes an unlimited number of arguments
// It's represented with the ... symbol

func main() {
	xi := []int{1, 2, 3, 4, 5}
	// Sum is asking for unlimited ints, not a slice.
	// Using the ... after the identifier means unfurling
	x := sum(xi...)
	fmt.Println("The total is", x)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
