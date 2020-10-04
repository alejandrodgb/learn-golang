package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}

	fmt.Println(p1)

	fmt.Println(p1.last)
}
