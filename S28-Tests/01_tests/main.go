package main

import "fmt"

func sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s = s + v
	}
	return s
}

func main() {
	fmt.Println(sum(2, 4, 5, 4, 3))
	fmt.Println(sum(5, 0, 8))
	fmt.Println(sum(5, 4))
}
