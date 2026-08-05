package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	got := Hello("Jorge")

	want := "Hello Jorge"

	require.Equal(t, got, want)
}
