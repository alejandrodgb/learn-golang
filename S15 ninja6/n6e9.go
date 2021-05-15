package main

import (
	"fmt"
)

func main() {
	floatStruct := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Sum:", sum(floatStruct...))
	fmt.Println("Divide:", divideSum(sum, floatStruct...))
	fmt.Println("Product:", divideProd(sum, floatStruct...))
}

func sum(x ...float64) float64 {
	sum := 0.0
	for _, v := range x {
		sum += v
	}
	return (sum)
}

func divideSum(f func(...float64) float64, x ...float64) float64 {
	sum := 0.0
	for _, v := range x {
		sum += 1.0 / float64(v)
	}
	return f(sum)
}

func divideProd(f func(...float64) float64, x ...float64) float64 {
	prod := 1.0
	for _, v := range x {
		prod *= 1.0 / float64(v)
	}
	return f(prod)
}
