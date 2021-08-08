package word

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alejandrodgb/learn-golang/S29-Ninja13/n13e2/quote"
)

func TestUseCount(t *testing.T) {
	m := UseCount("the dog that likes another dog eats the birds that eats birds")
	o := map[string]int{"the": 2, "dog": 2, "that": 2, "likes": 1, "another": 1, "eats": 2, "birds": 2}
	eq := reflect.DeepEqual(m, o)
	if !eq {
		t.Error("Expected", o, "Got", m)
	}
}

func TestCount(t *testing.T) {
	c := Count("Hello, this is a test")
	if c != 5 {
		t.Error("Expected:", 5, "Got:", c)
	}
}

func ExampleCount() {
	fmt.Println(Count("Hello world!"))
	// Output:
	// 2
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
