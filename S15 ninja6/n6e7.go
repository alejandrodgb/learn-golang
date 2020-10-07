package main

import "fmt"

func main() {
	f := funk
	fmt.Println(f())
}

func funk() string {
	return "Funky!"
}
