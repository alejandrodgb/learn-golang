// Package dog converts human years to dog years.
package dog

import "fmt"

// Years function takes in human years and returns dog years
func Years(h int) int {
	return h * 7
}

func main() {
	fmt.Println(Years(9))
}
