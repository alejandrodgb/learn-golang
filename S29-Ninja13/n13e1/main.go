package main

import (
	"fmt"

	"github.com/alejandrodgb/learn-golang/S29-Ninja13/n13e1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
