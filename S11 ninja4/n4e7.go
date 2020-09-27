package main

import "fmt"

func main() {
	x := [][]string{[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."}}

	for _, vi := range x {
		fmt.Println(vi)
		for _, vj := range vi {
			fmt.Println("\t", vj)
		}
	}
}
