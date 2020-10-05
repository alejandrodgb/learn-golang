package main

import "fmt"

func main() {
	anonymousStruct := struct {
		required  bool
		anonymous bool
		author    string
		length    int
	}{
		required:  true,
		anonymous: true,
		author:    "John",
		length:    5,
	}

	fmt.Println(anonymousStruct)

}
