package main

import (
	"fmt"
)

func main() {
	fmt.Println("Simple Bit Shifting")
	simpleBitShifting()
	fmt.Println("")
	fmt.Println("Iota Bit Shifting")
	iotaBitShifting()
}

func simpleBitShifting() {
	x := 2
	fmt.Printf("%d\t%b\n", x, x)

	// Shifting a bit on x to the left
	// if x = 2 (decimal) = 10 (binary)
	// then shifting a bit would make 100 (binary) = 2^2*1 + 2^1*0 + 2^0*0 = 4 (decimal)
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y)
}

const (
	_  = iota             // First iota is zero
	kb = 1 << (iota * 10) // This will shift by 10 which will give us 10000000000 (binary) or 1024 (decimal)
	mb = 1 << (iota * 10) // This will shift by 20 which will give us 100000000000000000000 (binary) or 1024*1024 (decimal)
	gb = 1 << (iota * 10) // This will shift by 30
)

func iotaBitShifting() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
