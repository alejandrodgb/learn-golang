// defer will defer the execution of a function until the exit point of the function that called it

package main

import "fmt"

func main() {
	nonDeferred()
	deferred()
}

func nonDeferred() {
	fmt.Println("Non Deferred")
	foo()
	bar()
}

func deferred() {
	fmt.Println("\nDeferred")
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}
