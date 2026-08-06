package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)

	exp := 5

	assert.Equal(t, exp, sum)
}

func ExampleAdd() {
	fmt.Println(Add(1, 4))
	// Output: 5
}
