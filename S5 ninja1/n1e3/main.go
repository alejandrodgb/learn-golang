// Using the code from the previous exercise,
// At the package level scope, assign the following values to the three variables
//   1. for x assign 42
//   2. for y assign “James Bond”
//   3. for z assign true
// in func main
//    1. use fmt.Sprintf to print all of the VALUES to one single string.
//      ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
//    2. print out the value stored by variable “s”

package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v, %v, %v", x, y, z)
	fmt.Println(s)
}
