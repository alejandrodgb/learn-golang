package main

import "fmt"

func main() {
	// Returning a string
	s1 := foo()
	fmt.Println("foo:", s1)

	// Returning a func
	x := bar()
	fmt.Printf("bar type: %T\n", x)
	fmt.Println("bar:", x())

}

func foo() string {
	return "Hello world"
}

// func() int is the type that will be returned
// in the return statement, it is the anonymous function of type func() int that is being returned
func bar() func() int {
	return func() int {
		return 451
	}

}
