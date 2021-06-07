// Write a program that:
// - puts 100 numbers to a channel
// - pull the numbers of the channel and print them

package main

import "fmt"

func main() {
	c := make(chan int)
	go write(c, 100)

	read(c)

	fmt.Println("Exiting program")
}

func write(c chan<- int, count int) {
	for i := 0; i < count; i++ {
		c <- i
	}
	close(c)
}

func read(c <-chan int) {
	for v := range c {
		fmt.Println("Reading from read():", v)
	}
}
