// Fix this code
// package main
//
// import (
// 	"fmt"
// )
// func main() {
// 	c := make(chan int)
//	c <- 42
//	fmt.Println(<-c)
// }

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Create an anonymous function as the channel will block and
	// if there is only the main goroutine then it will block that one
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
