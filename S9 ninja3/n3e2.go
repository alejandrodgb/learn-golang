package main

import "fmt"

func main() {
	start := 65
	end := 90
	for i := start; i <= end; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
