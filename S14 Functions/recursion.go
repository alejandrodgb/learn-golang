package main

import (
	"fmt"
)

func main() {
	fmt.Println("Factorial recursion:", factorialRecursion(4))
	fmt.Println("Factorial loop:", factorialLoop(4))
}

func factorialRecursion(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorialRecursion(x-1)
}

func factorialLoop(x int) int {
	for i := x - 1; i > 0; i-- {
		x *= i
	}
	return x
}
