package main

import "fmt"

func main() {
	rf := firstFunk()
	fmt.Println(rf())
	fmt.Printf("%T\n", rf)
}

func firstFunk() func() string {
	return func() string {
		return "This is a nested func."
	}
}
