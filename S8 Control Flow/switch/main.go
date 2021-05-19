package main

import (
	"fmt"
)

func main() {
	x := 10
	simpleSwitch(x)
	defaultSwitch(x)
	valueSwitch("Two")
}

func simpleSwitch(x int) {
	switch {
	case x < 5:
		fmt.Println("Case < 5")
	case x == 5:
		fmt.Println("Case == 5")
	case x > 5:
		fmt.Println("Case > 5")
	}
}

func defaultSwitch(x int) {
	switch {
	case x < 5:
		fmt.Println("Case < 5")
	case x == 5:
		fmt.Println("Case == 5")
	default:
		fmt.Println("Default switch")
	}
}

func valueSwitch(s string) {
	switch s {
	case "Name":
		fmt.Println("Name")
	case "String":
		fmt.Println("String")
	case "One", "Two", "Three":
		fmt.Println("Multiple values")
	default:
		fmt.Println("Default")
	}
}
