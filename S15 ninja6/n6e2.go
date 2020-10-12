package main

import (
	"fmt"
)

func main() {
	fmt.Println("foo:", foo([]int{1, 2, 3, 4, 5}...))
	fmt.Println("bar:", bar([]int{1, 2, 3, 4, 5}))
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}