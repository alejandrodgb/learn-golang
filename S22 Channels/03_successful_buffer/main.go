package main

import "fmt"

func main() {
	c := make(chan int, 1) // Creating a buffer allowing for 1 value to be passed to the channel
	// The channel will take it regardless whether something is ready to pull it off

	c <- 42

	fmt.Println(<-c)
}
