package main

import "fmt"

func send(v chan<- int) {
	v <- 10
	close(v)
}

func main() {
	c := make(chan int)

	go send(c)

	v, ok := <-c
	fmt.Println(v, ok)

	// This will fail if the channel is not closed because the reading operation will lock
	// and the main routine will not be able to continue - deadlock.
	v, ok = <-c
	fmt.Println(v, ok)
}
