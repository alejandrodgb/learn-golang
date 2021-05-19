package main

import "fmt"

func main() {
	p := person{
		name:     "James",
		lastname: "Bond",
		age:      33,
	}
	fmt.Println("Original:", p)
	changeMe(&p, "Jason", "Bourne", 22)
	fmt.Println("New:", p)
}

type person struct {
	name     string
	lastname string
	age      int
}

func changeMe(p *person, name string, lastname string, age int) {
	p.name = name
	p.lastname = lastname
	p.age = age
}
