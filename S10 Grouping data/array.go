// Arrays are not used. Use slices.
// The length is part of the type

package main

import (
	"fmt"
)

func main() {
	var x [5]int // Spefify data type and length of array.
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	// Length forms part of the type. Two arrays of differnet length are of different type.
	fmt.Printf("%T\n", x)
	var y [3]int
	fmt.Printf("%T\n", y)
}
