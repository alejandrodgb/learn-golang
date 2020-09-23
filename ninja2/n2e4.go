package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(100)
	fmt.Println("Original Number")
	fmt.Printf("%d, %b, %#x\n", x, x, x)

	y := x << 1
	fmt.Println("\nBit-Shifted Number")
	fmt.Printf("%d, %b, %#x\n", y, y, y)

}
