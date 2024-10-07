package iteration

import (
	"fmt"
	"testing"
)

// test for a function that repeats a character 5 times
func TestRepeat(t *testing.T) {
	t.Run("if iteration value is specified", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

// define a benchmark for Repeat function
func BenchmarkRepeat(b *testing.B) {
	// convention - to define the loop
	for i := 0; i < b.N; i++ {
		Repeat("a", 2)
	}
}

// document the function
func ExampleRepeat() {
	repeat := Repeat("a", 3)
	fmt.Println(repeat)
	// Output: aaa
}
