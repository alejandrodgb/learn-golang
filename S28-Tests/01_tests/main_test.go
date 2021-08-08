package main

import "testing"

func TestSum(t *testing.T) {
	x := sum(2, 4, 5, 4, 3)
	if x != 18 {
		t.Error("Expected", 18, "got", x)
	}
}
