package main

import (
	"fmt"
	"runtime"
	"sync"
)

func increment(counter *int, wg *sync.WaitGroup) {
	v := *counter
	runtime.Gosched()
	v++
	*counter = v
	wg.Done()
}

func main() {
	counter := 0
	const adds = 100

	var wg sync.WaitGroup
	wg.Add(adds)

	for i := 0; i < adds; i++ {
		go increment(&counter, &wg)
		//fmt.Println(counter)
	}

	wg.Wait()
	fmt.Println(counter)
}
