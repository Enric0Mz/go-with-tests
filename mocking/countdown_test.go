package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountdown(t *testing.T) {
	t.Run("Print 3", func(t *testing.T) {
		spySleeper := &SpyCoundDownOperations{}
		buffer := bytes.Buffer{}
		expect := `3
2
1
Go!`
		Countdown(&buffer, spySleeper)

		actual := buffer.String()

		require.Equal(t, expect, actual)
		require.Equal(t, 3, len(spySleeper.Calls))
	})
	t.Run("should sleep before every count", func(t *testing.T) {
		spySleeper := &SpyCoundDownOperations{}
		expect := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		Countdown(spySleeper, spySleeper)
		require.Equal(t, expect, spySleeper.Calls)
	})
}
