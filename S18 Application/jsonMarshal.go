package main

import (
	"encoding/json"
	"fmt"
)

// Person is a struct for person that takes
// first name, last name, and age
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   33,
	}
	p2 := Person{
		First: "Jason",
		Last:  "Bourne",
		Age:   44,
	}

	s := `[{"First":"James","Last":"Bond","Age":33},{"First":"Jason","Last":"Bourne","Age":44}]`

	people := []Person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nMarshaled data", string(bs))

	ubs := []byte(s)

	people2 := []Person{}
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(ubs, &people2)

	fmt.Println("\nUnmarshaled data", people2)

	for i, v := range people2 {
		fmt.Println("\n--- PERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
