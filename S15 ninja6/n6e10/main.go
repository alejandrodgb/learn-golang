package main

import (
	"fmt"
)

func main() {
	x := 1
	x++
	fmt.Println("Outside block x:", x)

	{
		y := 77.6
		y--
		fmt.Println("Inside block y:", y)
		fmt.Println("Inside block x:", x)
	}
	y := 10
	fmt.Println("Outside block y:", y)
}
