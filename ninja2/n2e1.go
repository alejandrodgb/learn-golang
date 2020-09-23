package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Intn(100)

	fmt.Printf("Decimal = %d\n", num)
	fmt.Printf("Binary = %b\n", num)
	fmt.Printf("Hex = %#x\n", num)
}
