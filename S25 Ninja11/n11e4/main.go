package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	r, err := sqrt(-9)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		se := sqrtError{
			lat:  "lat",
			long: "lon",
			err:  fmt.Errorf("Sqrt of negative number"),
		}
		return 0, se
	}
	return math.Sqrt(f), nil
}
