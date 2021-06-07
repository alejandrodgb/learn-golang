// Use comma,ok idiom starting from the code below
// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	c := make(chan int)
//
// 	v, ok :=
// 	fmt.Println(v, ok)
//
// 	close(c)
//
// 	v, ok =
// 	fmt.Println(v, ok)
// }

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(c chan<- int) {
		c <- 1000
	}(c)

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
