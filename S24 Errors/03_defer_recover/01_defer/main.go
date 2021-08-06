package main

import "fmt"

func main() {
	fmt.Println("A: ", a(0))
	b()
	fmt.Println("C: ", c(1))
	return
}

func a(i int) int {
	defer fmt.Println("A defer:", i)
	i++

	return i
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println("for range:", i)
	}
}
func c(i int) int {
	defer func() {
		fmt.Println("C ib4 defer", i)
		i++
		fmt.Println("C defer:", i)
	}()
	return 1
}
