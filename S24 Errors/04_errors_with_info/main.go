package main

import (
	"errors"
	"log"
	"math"
)

func main() {
	y := 9.0
	i, err := sqrt(y)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Sqrt of %v = %f", y, i)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Square root of negative number")
	}
	return math.Pow(f, 0.5), nil
}
