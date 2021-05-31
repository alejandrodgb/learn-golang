// The expression is a function and it can be assigned to a variable

package main

import "fmt"

func main() {

	// A func expression

	f := func() {
		fmt.Println("This is a func expression")
	}
	f()

	// A func expression with parameters
	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	g(1984)

}
