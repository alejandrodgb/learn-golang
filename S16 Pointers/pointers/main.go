package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // The & provides the address where the object is stored

	// The type will be different a is an int while &a is a pointer to an int
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// You can also assign the address to a value
	b := &a
	fmt.Println(b)

	// And we can get the value stored at an address with *
	fmt.Println(*b)

	// We can change the value by giving the value to the address

	*b = 43 // This will change a to 43
	fmt.Println(a)

}
