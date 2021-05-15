package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42 // This will fail because there is nothing ready to pull off the value and the channel will block.

	fmt.Println(<-c)
}
