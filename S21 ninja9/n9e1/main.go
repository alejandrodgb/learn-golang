package main

import (
	"fmt"
	"sync"
)

func multiply(wg *sync.WaitGroup, x ...float64) {
	fmt.Println("Enter multiply")
	var counter float64 = 1
	for _, value := range x {
		counter *= value
		//fmt.Println("Loop:", i, "\tCounter:", counter)
	}
	wg.Done()
	fmt.Println("Multiplying:", counter)

}

func sum(x ...float64) {
	fmt.Println("Enter sum")
	var counter float64
	for _, value := range x {
		counter += value
		//fmt.Println("Loop:", i, "\tCounter:", counter)
	}
	fmt.Println("Sum:", counter)
}

func main() {
	var wg sync.WaitGroup
	//fmt.Println("Routines", runtime.NumGoroutine())

	wg.Add(1)
	go multiply(&wg, 1, 2, 3, 4)
	//fmt.Println("Routines", runtime.NumGoroutine())

	sum(1, 2, 3, 4)
	//fmt.Println("Routines", runtime.NumGoroutine())

	wg.Wait()
}

// 50048371023 - electricity
// 50048371048 - water
