package iteration_test

import (
	"fmt"
	"testing"

	"callegarimattia.com/iteration"
)

func TestRepeat(t *testing.T) {
	t.Parallel()

	t.Run("repeate it 5 times",
		func(t *testing.T) {
			t.Parallel()

			repeated := iteration.Repeat("a", 5)
			expected := "aaaaa"

			if repeated != expected {
				t.Errorf("expected %s, got %s", expected, repeated)
			}
		})

	t.Run("repeate it 20 times",
		func(t *testing.T) {
			t.Parallel()

			repeated := iteration.Repeat("a", 20)
			expected := "aaaaaaaaaaaaaaaaaaaa"

			if repeated != expected {
				t.Errorf("expected %s, got %s", expected, repeated)
			}
		})
	t.Run("repeate it -1 times",
		func(t *testing.T) {
			t.Parallel()

			repeated := iteration.Repeat("aa", -1)
			expected := ""

			if expected != repeated {
				t.Errorf("expected %s, got %s", expected, repeated)
			}
		})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := iteration.Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
