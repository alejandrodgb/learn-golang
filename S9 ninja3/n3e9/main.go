package main

import "fmt"

func main() {

	favSport := "surf"

	switch favSport {
	case "motocross":
		fmt.Println("Motorcycles")
	case "surf":
		fmt.Println("Surfboards")
	default:
		fmt.Println("No sport")
	}

	boolTest()
}

func boolTest() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
