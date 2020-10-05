package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "John",
		last:     "Doe",
		icecream: []string{"vanilla", "chocolate"},
	}

	p2 := person{
		first:    "Sally",
		last:     "Fitzgibbon",
		icecream: []string{"coconut", "strawberry"},
	}

	fmt.Println(p1.first, "'s favourite icecream")
	for _, ic := range p1.icecream {
		fmt.Println("- ", ic)
	}

	fmt.Println(p2.first, "'s favourite icecream")
	for _, ic := range p2.icecream {
		fmt.Println("- ", ic)
	}
}
