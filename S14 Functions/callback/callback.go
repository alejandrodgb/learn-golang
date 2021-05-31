// A call back is passing a func as an argument

package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of all nubmers:", sum(ii...))
	fmt.Println("Sum of even numbers:", even(sum, ii...))
	fmt.Println("Sum of odd numbers:", odd(sum, ii...))
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, vi ...int) int {
	// This function will take a function and integers as parameters
	// and return the sum of all even numbers

	yi := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(x ...int) int, vi ...int) int {
	// This function will take a function and integers as parameters
	// and return the sum of all odd numbers

	yi := []int{}
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
