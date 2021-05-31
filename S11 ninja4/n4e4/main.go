package main

import "fmt"

func main() {
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	mySlice = append(mySlice, 52)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 53, 54, 55)
	fmt.Println(mySlice)

	y := []int{56, 57, 58, 59, 60}
	mySlice = append(mySlice, y...)
	fmt.Println(mySlice)
}
