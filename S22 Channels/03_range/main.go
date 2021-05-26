package main

import "fmt"

// Channels can have direction
// c<- is a receive-only channel
// <-c is a send-only channel
// Channels can be reassigned from general to specific but not the other way around: <-cs = c works, but c = <-cs doesn't

func foo(c chan<- int, v int) {
	for i := 0; i < 10; i++ {
		c <- i + 1
	}
	// Channel has to be closed here because of the scope of this variable. This is the channel where values are being
	// added.
	close(c)
}

func main() {

	c := make(chan int)

	// Send
	go foo(c, 52)

	// Receive
	// The range loop will continue to pull values from the channel until the channel is closed.
	// If there are no more values in the channel but the channel hasn't been closed it will
	// cause a deadlock.
	for v := range c {
		fmt.Println(v)
	}
	// If it is closed here, this is out of scope and deadlock happens. This is not the
	// channel where values are being added.
	// close(c)

	fmt.Println("About to exit")
}
