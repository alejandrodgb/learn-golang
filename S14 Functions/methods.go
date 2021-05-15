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
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
}
