package main

import (
	"fmt"
	"sync"
)

func increment(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	v := *counter
	v++
	*counter = v
	mu.Unlock()
	wg.Done()
}

func main() {
	counter := 0
	const adds = 100

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(adds)

	for i := 0; i < adds; i++ {
		go increment(&counter, &wg, &mu)
		//fmt.Println(counter)
	}

	wg.Wait()
	fmt.Println(counter)
}
