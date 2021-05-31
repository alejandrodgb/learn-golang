package main

import "fmt"

func main() {
	name := "Johna"
	if name == "John" {
		fmt.Println("John Doe")
	} else if name == "Jane" {
		fmt.Println("Jane Doe")
	} else {
		fmt.Println("No Name")
	}
}
