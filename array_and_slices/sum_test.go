package sum

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()

	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	assert.Equal(t, want, got)
}

func TestSumAll(t *testing.T) {
	t.Parallel()

	got := All([][]int{{1, 2}, {2}})
	want := []int{3, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	sliceInputs := GetInputs(b)
	for i := 0; i < b.N; i++ {
		All(sliceInputs)
	}
}

func GetInputs(b *testing.B) [][]int {
	b.Helper()

	random := [1000][1000]int{}

	for row := 0; row < 1000; row++ {
		for col := 0; col < 1000; col++ {
			val, err := rand.Int(rand.Reader, big.NewInt(10000))
			if err != nil {
				val = big.NewInt(0)
			}

			random[row][col] = int(val.Int64())
		}
	}

	var inputs [][]int

	for i := 0; i < len(inputs); i++ {
		// Each row is a slice
		inputs = append(inputs, random[i][:])
	}

	return inputs
}
