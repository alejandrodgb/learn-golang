package main

import "fmt"

func unsuccessfulChannel() {
	// CHANNELS BLOCK: channels block until the load and unload can happen at the same time.
	// This function will not run because when the channel gets loaded it blocks and stops the
	// code from continuing to run. The function cannot finalise
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}

func successfulChannel() {
	// This will run because the channel will block the anonymous function but the remainder of the
	// code will continue to run. Once the value is pulled off of the channel, then the anonymous
	// function will complete
	c := make(chan string)
	go func() {
		c <- "successfulChannel"
	}()
	fmt.Println(<-c)
}

func successfulBuffer() {
	// A buffer channel will allow values to be sitting in the channel
	// The code will run because once the value gets to the channel, there will be space for it to be in the channel
	// and sit in the channel. Once the fmt is ready to pull it, the value will be ready to go.
	c := make(chan string, 1)
	c <- "successfulBuffer"
	fmt.Println(<-c)
}

func unsuccessfulBuffer() {
	// This will not run because the channel only has space to store one value. Once the value is put in there,
	// the next value will block until unloaded but it won't be able to unload because the routine is locked and
	// will not continue to advance.
	c := make(chan int, 1)
	c <- 44
	c <- 45
	fmt.Println(<-c)
}

func successfulMultipleBuffer() {
	// This will not run because the channel only has space to store one value. Once the value is put in there,
	// the next value will block until unloaded but it won't be able to unload because the routine is locked and
	// will not continue to advance.
	c := make(chan string, 2)
	c <- "successfulMultipleBuffer1"
	c <- "successfulMultipleBuffer2"
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func main() {
	// unsuccessfulChannel()
	successfulChannel()
	successfulBuffer()
	// unsuccessfulBuffer()
	successfulMultipleBuffer()
}
