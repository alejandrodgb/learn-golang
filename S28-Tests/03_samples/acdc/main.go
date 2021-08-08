// Package acdc asks if you are ready to rock
package acdc

// Sum sums over a given number of integers
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s = s + v
	}
	return s
}
