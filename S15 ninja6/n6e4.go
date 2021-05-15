package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and I am %v years old\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "John",
		last:  "Smith",
		age:   33,
	}

	p1.speak()
}
