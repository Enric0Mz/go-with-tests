package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	t.Run("should return Hello name when name is provided", func(t *testing.T) {
		got := Hello("Jorge")
		want := "Hello Jorge"

		require.Equal(t, got, want)
	})
	t.Run("should return Hello Stranger when no name is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello Stranger"

		require.Equal(t, got, want)
	})

}
