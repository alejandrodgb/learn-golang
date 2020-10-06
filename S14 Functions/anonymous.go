package main

import "fmt"

func main() {
	// A normal function
	foo()

	// An anonymous function
	func() {
		fmt.Println("Anonymous function")
	}()

	// An anonymous function with parameters
	func(x string) {
		fmt.Println("Anonymous func with parameter:", x)
	}("This is parameter x")
}

func foo() {
	fmt.Println("Foo func")
}
