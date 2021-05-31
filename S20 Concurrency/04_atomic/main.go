// Mutex locks a variable that is being accesses to avoid a race condition

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type results struct {
	min int
	max int
}

func minMax(v []int) (r results) {

	min := v[0]
	max := v[0]

	// Find minimum value
	for i, e := range v {
		if i == 0 || e < min {
			min = e
		}
	}

	// Find maximum falue
	for _, e := range v {
		if e > max {
			max = e
		}
	}

	return results{
		min: min,
		max: max,
	}
}

func main() {

	fmt.Println("CPUs:          ", runtime.NumCPU())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup

	var goRoutines []int

	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {

			atomic.AddInt64(&counter, 1)

			runtime.Gosched()

			// Value has to be read with atomic because it is in a goroutine and
			// reading it independently may cause a race condition because other
			// routines may also be using that value
			fmt.Println("Goroutine counter: ", atomic.LoadInt64(&counter))

			wg.Done()
		}()
		//fmt.Println("Go routines:", runtime.NumGoroutine())
		goRoutines = append(goRoutines, runtime.NumGoroutine())
	}
	wg.Wait()

	// Counter can be read here without race condition as the waitgroup has finalised
	// and there will be no more goroutines running
	fmt.Println("Counter:       ", counter)
	//fmt.Println("Max No. Goroutines:", goRoutines)
	fmt.Println("Min Goroutines:", minMax(goRoutines).min)
	fmt.Println("Max Goroutines:", minMax(goRoutines).max)
}
