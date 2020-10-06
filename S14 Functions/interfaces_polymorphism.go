package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// Structure of a function:
//   func (r receiver) identifier(paremeters) (return(s)) { ... }
// The receiver attaches the function to the type so the function has access to the type

func (s secretAgent) speak() {
	fmt.Println("Secret agent:", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("Person:	", p.first, p.last)
}

// This will make any value with the method speak ALSO of type human
// A value can be of more than one type
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Bar - I am a person", h.(person).first)
	case secretAgent:
		fmt.Println("Bar - I am a secret agent", h.(secretAgent).first)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	p1 := person{
		first: "Enrico",
		last:  "Fermi",
	}

	fmt.Println(sa1)
	sa1.speak()

	bar(p1)
	bar(sa1)
}
