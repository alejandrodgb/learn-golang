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
	const totalGoroutines = 10
	const counter = 10
	var wg sync.WaitGroup
	c := make(chan int)

	go createGoroutines(totalGoroutines, counter, c, &wg)

	wg.Wait()
	close(c)

	for v := range c {
		fmt.Println(v)
	}

}

func createGoroutines(totalGoroutines, counter int, c chan<- int, wg *sync.WaitGroup) {
	fmt.Println("Create goroutines")
	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func(s, n int) {
			for i := s * 10; i < s*10+n; i++ {
				c <- i
			}
			wg.Done()
		}(i, counter)
	}
}

func write(s, n int, c chan<- int, wg *sync.WaitGroup) {
	for i := s * 10; i < s*10+n; i++ {
		c <- i
	}
	wg.Done()
}

//67858
