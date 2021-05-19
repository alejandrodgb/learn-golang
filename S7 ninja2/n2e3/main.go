package main

import "fmt"

const (
	tc float64 = 12.2
	uc         = "Salsa"
)

func main() {
	fmt.Printf("TYPED constant = %v\n", tc)
	fmt.Printf("UNTYPED constant = %v\n", uc)
}
