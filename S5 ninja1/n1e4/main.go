// For this exercise
//    1. Create your own type. Have the underlying type be an int.
//    2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
//    3. in func main
//      3.1 print out the value of the variable “x”
//      3.2 print out the type of the variable “x”
//      3.3 assign 42 to the VARIABLE “x” using the “=” OPERATOR
//      3.4 print out the value of the variable “x”

package main

import "fmt"

type taco int

var x taco

func main() {
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)

}
