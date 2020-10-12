package main

import (
	"fmt"
	"sort"
	"strings"
)

// Min returns the min of two int values
func Min(x1, x2 int) int {
	if x1 > x2 {
		return x2
	}
	return x1
}

type Person struct {
	First string
	Age   int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

// Len is the number of elements in the collection.
func (a ByAge) Len() int {
	return len(a)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// Swap swaps the elements with indexes i and j.
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByFirst implements sort.Interface for []Person based
// on the First field.
type ByFirst []Person

// Len is the number of elements in the collection.
func (f ByFirst) Len() int {
	return len(f)
}

// Less reports whether the element with
// index i shoudl sort before the element with index j.
func (f ByFirst) Less(i, j int) bool {
	// Convert to lower case all strings
	si := strings.ToLower(f[i].First)
	sj := strings.ToLower(f[j].First)

	xi := len(si)
	xj := len(sj)

	// Range through the strings up to the len of the
	// shortest string
	for xsi := 0; xsi < Min(xi, xj); xsi++ {

		if si[xsi] < sj[xsi] {
			return true
		} else if si[xsi] > sj[xsi] {
			return false
		}
	}
	// If all letters are the same, order them by
	// the length of the string
	if xi < xj {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (f ByFirst) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("People:\n", people)
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by Age People\n", people)

	sort.Sort(ByFirst(people))
	fmt.Println("Sorted by First People\n", people)

}
