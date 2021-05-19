// Set of methods attached to a type
// A non-pointer receiver will work with both pointer and non-pointer values
// A pointer value will only work with pointer values

package main

import (
	"fmt"
	"math"
)

func main() {
	circleShape := circle{
		radius: 5,
	}

	//fmt.Println("Non-pointer value, non-pointer receiver:", info(circleShape))  // Non-pointer value. This will fail if there is a pointer receiver.
	fmt.Println("Pointer value, non-pointer receiver:    ", info(&circleShape)) //Pointer value
	fmt.Println("Pointer value, pointer receiver", info(&circleShape))          //Pointer value
}

type circle struct {
	radius float64
}

type results struct {
	area      float64
	perimeter float64
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
	perimeter() float64
}

func info(s shape) results {
	return results{
		area:      s.area(),
		perimeter: s.perimeter(),
	}
}
