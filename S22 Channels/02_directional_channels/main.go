package main

import "fmt"

// Channels can have direction
// c<- is a receive-only channel
// <-c is a send-only channel
// Channels can be reassigned from general to specific but not the other way around: <-cs = c works, but c = <-cs doesn't

func foo(c chan<- int, v int) {
	c <- v
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

func main() {

	c := make(chan int)

	// Send
	go foo(c, 52)

	// If this function is sent as a goroutine the process will do nothing because the fmt
	// statement will not wait for foo and bar to finish. If it does not have a goroutine,
	// then the channel will block until it receives information which will allow all code to run
	bar(c)

	fmt.Println("About to exit")
}
