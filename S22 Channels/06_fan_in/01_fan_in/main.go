package main

import (
	"fmt"
	"sync"
)

type results struct {
	channel string
	number  int
}

func send(e, o chan<- int, m int) {
	for i := 0; i <= m; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive(e, o <-chan int, fanin chan<- results) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanin <- results{"even", v}
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- results{"odd", v}
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}

func main() {
	e := make(chan int)
	o := make(chan int)
	f := make(chan results)

	go send(e, o, 10)

	go receive(e, o, f)

	for v := range f {
		fmt.Println(v)
	}
	fmt.Println("Exiting program")
}
