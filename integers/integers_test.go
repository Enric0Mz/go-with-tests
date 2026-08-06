package integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)

	exp := 5

	assert.Equal(t, exp, sum)
}
