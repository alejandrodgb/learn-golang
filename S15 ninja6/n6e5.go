package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	squareShape := square{
		side: 15,
	}
	circleShape := circle{
		radius: 12.345,
	}
	info(squareShape)
	info(circleShape)
}
