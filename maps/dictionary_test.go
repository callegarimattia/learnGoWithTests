package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	t.Run("known word", func(t *testing.T) {
		t.Parallel()

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assert.Equal(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		t.Parallel()

		got, err := dictionary.Search("unknown")
		assert.Equal(t, "", got)
		assert.ErrorIs(t, err, errNotFound)
	})
}
