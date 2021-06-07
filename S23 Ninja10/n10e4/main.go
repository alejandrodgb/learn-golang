// Starting with the code below, pull values off the channel using a select statement
// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	q := make(chan int)
// 	c := gen(q)
//
// 	receive(c, q)
//
// 	fmt.Println("about to exit")
// }
//
// func gen(q <-chan int) <-chan int {
// 	c := make(chan int)
//
// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}
//
// 	return c
// }

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		//q <- 1
		close(q)
		// If q is not closed or a value loaded into it the  receive() infinite
		// loop will go on forever because <-c will continue to pass value 0 as it is closed
	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("From channel c", v)
		case v := <-q:
			fmt.Println("Quitting", v)
			return
		}
	}
}
