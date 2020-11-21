package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
	lang  string
}

func (p *person) speak() {
	fmt.Println("Person", p.first, " speaks", p.lang)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "John",
		last:  "Doe",
		age:   33,
		lang:  "English",
	}

	saySomething(&p1)
}
