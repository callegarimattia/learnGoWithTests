package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d, got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
