package main

import "fmt"

func main() {
	fmt.Println(func() string {
		return "This is an anonymous func"
	}())
}
