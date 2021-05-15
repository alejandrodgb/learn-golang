package main

import "fmt"

// Send channel
func send(count int, even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// Receive channel
func receive(count int, even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("[even]:", v)
		case v := <-odd:
			fmt.Println("[odd]:", v)
		case i, ok := <-quit:
			fmt.Println("!ok", i, ok)
			return
		}
	}
}

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	const count int = 10

	// Send
	go send(count, even, odd, quit)

	// Receive
	receive(count, even, odd, quit)

	fmt.Println("Exit")
}
