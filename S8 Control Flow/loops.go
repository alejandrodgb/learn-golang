package main

import "fmt"

const (
	ol = 10
	il = 5
)

func main() {
	//loopIntro()
	//nestingLoops()
	//simpleFor()
	//eternalLoop()
	breakContinue()
}

func loopIntro() {
	// Loops follow an init; condition; post {} structure
	// Loop to count from 0 to 100

	for i := 0; i <= ol; i++ {
		fmt.Println(i)
	}
}

func nestingLoops() {
	for i := 0; i <= ol; i++ {
		fmt.Printf("Outer loop: %d\n", i)
		for j := 0; j <= il; j++ {
			fmt.Printf("\t Inner loop: %d\n", j)
		}
	}
}

func simpleFor() {
	a := 1
	for a < il {
		fmt.Println(a)
		a++
	}
}

func eternalLoop() {
	a := 1
	for {
		if a > il {
			break
		}
		fmt.Println(a)
		a++
	}
}

func breakContinue() {
	x := -1
	for {
		x++
		if x > ol {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
