package main

import (
	"log"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		log.Println("foo: ", i)
	}
	// Indicate to the waitgroup that the subroutine is done
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		log.Println("bar: ", i)
	}
}

func main() {
	log.Println("OS:           ", runtime.GOOS)
	log.Println("Arch:         ", runtime.GOARCH)
	log.Println("CPUs          ", runtime.NumCPU())
	log.Println("Goroutines b4:", runtime.NumGoroutine())

	// Add a waitgroup so foo() will execute
	wg.Add(1)

	// Launching a go routine
	go foo()

	log.Println("Goroutines after:", runtime.NumGoroutine())
	// Foo will not be seen in the output because after the goroutine is created the program continues with its
	// normal execution and finishes before foo is printed.
	// To get around this we add a waitgroup see https://pkg.go.dev/sync#WaitGroup
	bar()

	// Wait for all goroutines to execute before exiting
	wg.Wait()
}
