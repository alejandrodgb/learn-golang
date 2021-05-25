package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hola")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Adam", 33}
	saySomething(&p1)
}
