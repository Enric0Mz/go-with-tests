package iteration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	exp := "aaaaa"

	require.Equal(t, got, exp)
}

func BenchmarkLoop(b *testing.B) {
	for b.Loop() {
		Repeat("a", 50)
	}
}
