package main

import "fmt"

func main() {
	defer deferredFunction()
	normalFunction("John Smith")
}

func deferredFunction() {
	fmt.Println("This is the deferred function")
}

func normalFunction(s string) {
	fmt.Println("Hello", s)
}
