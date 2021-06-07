// Write a program that
// - launches 10 goroutines
//   - each goroutine adds 10 numbers to a channel
// - pull the numbers off the channel and print them

package main

import "fmt"

func main() {
	c := make(chan int)
	const totalGoroutines = 10
	const counter = 10

	for i := 0; i < totalGoroutines; i++ {
		go write(i, counter, c)
	}

	for i := 0; i < totalGoroutines*counter; i++ {
		fmt.Println(<-c)
	}

}

func write(s, n int, c chan<- int) {
	for i := s * 10; i < s*10+n; i++ {
		c <- i
	}
}
