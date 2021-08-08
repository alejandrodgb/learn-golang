// Package word has functions to help manage words
package word

import (
	"strings"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

// UseCount creates a map of a given string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++

	}
	return m
}

// Count is a function to count string
func Count(s string) int {
	return len(strings.Fields(s))
}
