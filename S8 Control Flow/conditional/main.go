package main

import (
	"fmt"
)

func main() {
	initStatement()
	ifElse()
}

func initStatement() {
	// Conditionals can have an initialization statement
	if x := 42; x == 42 {
		fmt.Println(x)
	}
}

func ifElse() {
	x := 41
	if x == 40 {
		fmt.Println("Value was 40")
	} else if x == 41 {
		fmt.Println("Value was 41")
	} else {
		fmt.Println("Value was not 40 or 41")
	}
}
