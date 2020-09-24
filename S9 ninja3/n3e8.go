package main

import "fmt"

func main() {

	num := 2

	switch {
	case num == 1:
		fmt.Println(1)
	case num == 2:
		fmt.Println(2)
	default:
		fmt.Println("No Number")
	}
}
