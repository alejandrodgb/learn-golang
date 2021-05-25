// Hands-on exercise #5
// Fix the race condition you created in exercise #3 by using package atomic

package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

var incrementer int64
var wg sync.WaitGroup

func increment() {
	atomic.AddInt64(&incrementer, 1)
	log.Println(atomic.LoadInt64(&incrementer))
	wg.Done()
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println(incrementer)
}
