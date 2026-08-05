package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	t.Run("should return Hello name when name is provided", func(t *testing.T) {
		got := Hello("Jorge", "")
		want := "Hello Jorge"

		require.Equal(t, want, got)
	})
	t.Run("should return Hello Stranger when no name is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello Stranger"

		require.Equal(t, want, got)
	})

	t.Run("should return in spanish when language is provided", func(t *testing.T) {
		got := Hello("Clark", "Spanish")
		want := "Hola Clark"

		require.Equal(t, want, got)

	})
	t.Run("should return in french when language is provided", func(t *testing.T) {
		got := Hello("Clark", "French")
		want := "Bonjour Clark"

		require.Equal(t, want, got)

	})

}
