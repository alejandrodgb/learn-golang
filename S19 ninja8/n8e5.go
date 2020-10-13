package main

import (
	"fmt"
	"sort"
)

// User ...
type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []User

// Len ...
func (a ByAge) Len() int {
	return len(a)
}

// Swap ...
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less ...
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// ByLast implements sort.Interface for []User based on
// the Last field.
type ByLast []User

// Len ...
func (l ByLast) Len() int {
	return len(l)
}

// Swap ...
func (l ByLast) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Less ...
func (l ByLast) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func printField(user []User, field string, sortType bool, title string) {
	fmt.Println("\n", title)
	switch field {
	case "full":
		for i, v := range user {
			fmt.Println(i, "| Age:", v.Age, "| Last:", v.Last)

			if sortType {
				sort.Strings(v.Sayings)
			}

			for _, vv := range v.Sayings {
				fmt.Println("-", vv)
			}
			fmt.Println()
		}
	case "summary":
		for i, v := range user {
			fmt.Println(i, "| Age:", v.Age, "| Last:", v.Last)
		}
	}
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	printField(users, "summary", false, "Original")
	sort.Sort(ByAge(users))
	printField(users, "full", false, "ByAge")
	sort.Sort(ByLast(users))
	printField(users, "full", true, "ByLast")
}
