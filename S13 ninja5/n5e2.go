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

	people := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range people {
		fmt.Println(v.first, v.last)
		for _, vv := range v.icecream {
			fmt.Println("- ", vv)
		}
		fmt.Println("----------")
	}
}
