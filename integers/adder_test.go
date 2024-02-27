package integers_test

import (
	"fmt"
	"testing"

	"callegarimattia.com/integers"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	sum := integers.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d, got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
