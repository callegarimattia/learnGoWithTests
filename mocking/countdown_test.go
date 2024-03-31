package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &spySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!
`

	require.Equal(t, want, got)
	require.Equal(t, countdownStart, spySleeper.Calls)
}
