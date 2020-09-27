package main

import "fmt"

func main() {
	mySlice := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", mySlice)

}
