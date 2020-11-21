package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func increment(counter *int64, wg *sync.WaitGroup) {
	atomic.AddInt64(counter, 1)
	//fmt.Println("Counter:", atomic.LoadInt64(counter))
	wg.Done()
}

func main() {
	var counter int64 = 0
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
