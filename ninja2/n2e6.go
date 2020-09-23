package main

import "fmt"

const (
	_   = iota
	yp1 = 2020 + iota
	yp2
	yp3
	yp4
)

func main() {
	fmt.Println(yp1)
	fmt.Println(yp2)
	fmt.Println(yp3)
	fmt.Println(yp4)
}
