// Mutex locks a variable that is being accesses to avoid a race condition

package main

import (
	"fmt"
	"runtime"
	"sync"
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

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	var goRoutines []int

	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			mu.Lock() // Lock the counter variable
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // Unlock the counter variable
			wg.Done()
		}()
		//fmt.Println("Go routines:", runtime.NumGoroutine())
		goRoutines = append(goRoutines, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter:       ", counter)
	//fmt.Println("Max No. Goroutines:", goRoutines)
	fmt.Println("Min Goroutines:", minMax(goRoutines).min)
	fmt.Println("Max Goroutines:", minMax(goRoutines).max)
}
