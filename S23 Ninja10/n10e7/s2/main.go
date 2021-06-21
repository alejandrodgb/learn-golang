// Write a program that
// - launches 10 goroutines
//   - each goroutine adds 10 numbers to a channel
// - pull the numbers off the channel and print them

package main

import (
	"fmt"
	"sync"
)

func main() {
	const totalGoroutines = 2
	const counter = 10
	var wg sync.WaitGroup
	c := make(chan int)

	go createGoroutines(totalGoroutines, counter, c, &wg)

	for v := range c {
		fmt.Printf("Read %d from c\n", v)
	}
}

func createGoroutines(totalGoroutines, counter int, c chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go write(i, counter, c, wg)
	}
	wg.Wait()
	close(c)
}

func write(s, n int, c chan<- int, wg *sync.WaitGroup) {
	for i := s * 10; i < s*10+n; i++ {
		fmt.Printf("Loaded %d into c\n", i)
		c <- i
	}
	wg.Done()
}
