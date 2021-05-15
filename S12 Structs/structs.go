package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   30,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.last)
	fmt.Println(p2.first)
}
