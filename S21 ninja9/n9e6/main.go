// Hands-on exercise #6
// Create a program that prints out your OS and ARCH. Use the following commands to run it
// go run
// go build
// go install - this will install under go/bin/n9e6

package main

import (
	"log"
	"runtime"
)

func main() {
	log.Println("OS:", runtime.GOOS)
	log.Println("ARCH:", runtime.GOARCH)
}
