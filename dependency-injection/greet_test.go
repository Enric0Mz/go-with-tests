package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	expected := "Hello, Chris"
	actual := buffer.String()

	require.Equal(t, expected, actual)

}
