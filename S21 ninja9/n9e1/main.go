package main

import (
	"fmt"
	"log"
	"sync"
)

var wg sync.WaitGroup

func sum(l []int) (int, error) {
	if len(l) == 0 {
		return 0, fmt.Errorf("Empty list")
	}
	var result int
	for _, v := range l {
		result = result + v
	}
	log.Printf("Sum = %d for list %d", result, l)
	wg.Done()
	return result, nil
}

func product(l []int) (int, error) {
	if len(l) == 0 {
		return 0, fmt.Errorf("Empty list")
	}
	var result int = 1
	for _, v := range l {
		result = result * v
	}
	log.Printf("Product = %d for list %d", result, l)
	wg.Done()
	return result, nil
}

func sumProduct(l []int, w []int) (int, error) {
	if len(l) != len(w) {
		return 0, fmt.Errorf("Lists must be equal length")
	} else if len(l) == 0 {
		return 0, fmt.Errorf("Empty list")
	}
	var result int
	for i := 0; i < len(l); i++ {
		temp := l[i] * w[i]
		result = result + temp
	}
	log.Printf("sumProduct = %d for list %d and weights %d", result, l, w)
	wg.Done()
	return result, nil
}

func main() {

	list := []int{}

	for i := 1; i < 10; i++ {
		wg.Add(3)
		list = append(list, i)
		//fmt.Println(list)
		go sum(list)
		go product(list)
		go sumProduct(list, list)
	}
	wg.Wait()
}
