package main

import "fmt"

func main() {
	var myArray [5]int
	for i := range myArray {
		myArray[i] = i + 100
	}

	for i, v := range myArray {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", myArray)
}
