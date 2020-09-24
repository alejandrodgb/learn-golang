package main

import "fmt"

func main() {
	yob := 1989
	for {
		if yob > 2020 {
			break
		}
		fmt.Println(yob)
		yob++
	}
}
