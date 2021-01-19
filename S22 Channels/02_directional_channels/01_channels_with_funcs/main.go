package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)
	// bar function cannot be run as a goroutine for the code to work.
	// bar will block until it receives from the channel so the code will be
	// finished. If a goroutine is sent then both will go do their thing
	// on their own and the code will not complete.

	fmt.Println("About to exit")
}

// send
func foo(c chan<- int) {
	fmt.Println("[foo]")
	c <- 11
}

// receive
func bar(c <-chan int) {
	fmt.Println("[bar]")
	fmt.Println(<-c)
}
