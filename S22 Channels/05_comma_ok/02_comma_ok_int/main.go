package main

import "fmt"

func send(e, o, q chan<- int) {
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the eve channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("From the quit channel", v, ok)
				return
			}
			fmt.Println("From the quit channel:", v)
			return
		}
	}
}

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go send(eve, odd, quit)

	// Receive
	receive(eve, odd, quit)

	fmt.Println("About to exit")
}
