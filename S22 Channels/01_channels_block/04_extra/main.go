package main

import (
	"fmt"
)

func squares(c chan int) {
	fmt.Println("[squres] reading")
	num := <-c
	c <- num * num
}

func cubes(c chan int) {
	fmt.Println("[cubes] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("[main] started")

	// Create channels
	squaresChan := make(chan int)
	cubesChan := make(chan int)

	// Start goroutines
	go squares(squaresChan)
	go cubes(cubesChan)

	testNum := 3

	fmt.Println("[main] sent testNum to squaresChan")

	squaresChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubesChan")

	cubesChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	squareVal, cubeVal := <-squaresChan, <-cubesChan
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of squareVal and cubeVal = ", sum)
	fmt.Println("[main] finished")

}
