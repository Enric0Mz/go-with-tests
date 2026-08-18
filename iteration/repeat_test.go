package iteration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	exp := "aaaaa"

	require.Equal(t, got, exp)
}
