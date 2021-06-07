// Starting with the code below pull the values off
// the channel using a for range loop

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	c := gen()
// 	receive(c)
//
// 	fmt.Println("about to exit")
// }
//
// func gen() <-chan int {
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
	c := gen() //If gen() is not modified and a go routine added then it will block main
	receive(c)

	fmt.Println("About to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
