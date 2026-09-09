package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountdown(t *testing.T) {
	t.Run("Print 3", func(t *testing.T) {
		spySleeper := &SpySleeper{}
		buffer := bytes.Buffer{}
		expect := `3
2
1
Go!`
		Countdown(&buffer, spySleeper)

		actual := buffer.String()

		require.Equal(t, expect, actual)
		require.Equal(t, 3, spySleeper.Calls)
	})
}
