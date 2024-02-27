package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Parallel()

	t.Run("hello world",
		func(t *testing.T) {
			t.Parallel()

			got := Hello("", "")
			want := "Hello, world!"
			assertCorrect(t, got, want)
		},
	)
	t.Run("hello Mario",
		func(t *testing.T) {
			t.Parallel()

			got := Hello("Mario", "")
			want := "Hello, Mario!"
			assertCorrect(t, got, want)
		})
	t.Run("in spanish",
		func(t *testing.T) {
			t.Parallel()

			got := Hello("Elodie", "Spanish")
			want := "Hola, Elodie!"
			assertCorrect(t, got, want)
		})
	t.Run("in french",
		func(t *testing.T) {
			t.Parallel()

			got := Hello("Luca", "French")
			want := "Bonjour, Luca!"
			assertCorrect(t, got, want)
		})
	t.Run("in italian",
		func(t *testing.T) {
			t.Parallel()

			got := Hello("Mattia", "Italian")
			want := "Buongiorno, Mattia!"
			assertCorrect(t, got, want)
		})
	t.Run("in portuguese",
		func(t *testing.T) {
			t.Parallel()

			got := Hello("Elodie", "Portuguese")
			want := "Oi, Elodie!"
			assertCorrect(t, got, want)
		})
}

func assertCorrect(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
