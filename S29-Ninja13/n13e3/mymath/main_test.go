package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		result float64
	}

	tests := []test{
		{
			data:   []int{1, 2, 3, 4, 6},
			result: 3,
		},
		{
			data:   []int{1, 4, 6, 40},
			result: 5,
		},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.result {
			t.Error("Expected", v.result, "Got", x)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 40}))
	//Output:
	// 5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{2, 3, 4, 55, 6, 5})
	}
}
