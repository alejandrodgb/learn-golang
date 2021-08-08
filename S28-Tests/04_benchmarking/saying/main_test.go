package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Alex")
	if s != "Welcome my dear Alex" {
		t.Error("Got", s, "Expected", "Welcome my dearest Alex")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Alex"))
	//Output:
	//Welcome my dear Alex
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Alex")
	}
}
