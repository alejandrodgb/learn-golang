package main

import "testing"

func TestSum(t *testing.T) {

	type test struct {
		data     []int
		result   int
		testname string
	}

	tests := []test{
		{data: []int{1, 2, 3}, result: 6, testname: "Test 1"},
		{data: []int{1, 2}, result: 3, testname: "Test 2"},
		{data: []int{1, 3}, result: 3, testname: "Test 3"},
		{data: []int{-1, -2, 3}, result: 0, testname: "Test 4"},
		{data: []int{-1, -2, -3}, result: -6, testname: "Test 5"},
	}

	for _, v := range tests {
		x := sum(v.data...)
		if x != v.result {
			t.Errorf("Test: %v Expected %v, got %v", v.testname, v.result, x)
		}
	}
}
