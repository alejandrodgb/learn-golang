package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("About to exit")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // This will close the channel in foo after it has sent 100 values (0-99)
	// to it. If it is not closed, then bar will deadlock because it will continue
	// to wait for a value but the channel has nothing else in it.
}

// receive
func bar(c <-chan int) {
	for v := range c {
		// Ranging over a channel will look at all values in the channel until that channel
		// has been closed.
		fmt.Println(v)
	}
}
