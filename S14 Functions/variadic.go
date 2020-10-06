package main

import "fmt"

// A variadic parameter is one that takes an unlimited number of arguments
// It's represented with the ... symbol

func main() {
	x := sum(1, 2, 3, 4, 5)
	fmt.Println("The total is", x)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
