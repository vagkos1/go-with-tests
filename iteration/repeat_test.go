package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat k times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
