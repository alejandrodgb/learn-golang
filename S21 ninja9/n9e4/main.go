// Hands-on exercise #4
// Fix the race condition you created in the previous exercise by using a mutex
// it makes sense to remove runtime.Gosched()

package main

import (
	"fmt"
	"sync"
)

var incrementer int
var wg sync.WaitGroup
var mu sync.Mutex

func increment() {
	mu.Lock()
	v := incrementer
	v = v + 1
	incrementer = v
	mu.Unlock()
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
