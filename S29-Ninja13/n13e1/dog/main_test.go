package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {

	type test struct {
		data     int
		result   int
		testname string
	}

	tests := []test{
		{
			data:     2,
			result:   14,
			testname: "YearsT1",
		},
		{
			data:     5,
			result:   35,
			testname: "YearsT2",
		},
	}

	for _, v := range tests {
		x := Years(v.data)
		if x != v.result {
			t.Errorf("Test: %v Expected %v, got %v", v.testname, v.result, x)
		}
	}
}

func TestYearsTwo(t *testing.T) {

	type test struct {
		data     int
		result   int
		testname string
	}

	tests := []test{
		{
			data:     2,
			result:   14,
			testname: "YearsTwoT1",
		},
		{
			data:     5,
			result:   35,
			testname: "YearsTwoT2",
		},
	}

	for _, v := range tests {
		x := YearsTwo(v.data)
		if x != v.result {
			t.Errorf("Test: %v Expected %v, got %v", v.testname, v.result, x)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
