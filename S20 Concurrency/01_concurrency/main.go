package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OOS:         ", runtime.GOOS)
	fmt.Println("ARCH:        ", runtime.GOARCH)
	fmt.Println("No. CPU:     ", runtime.NumCPU())
	fmt.Println("No. Routines:", runtime.NumGoroutine())

	wg.Add(1)
	go foo() // Keyword go starts a Go routine
	bar()

	fmt.Println("No. CPU:     ", runtime.NumCPU())
	fmt.Println("No. Routines:", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
