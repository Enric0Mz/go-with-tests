package structsmethodsinterfaces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerimeter(t *testing.T) {
	exp := 40.0

	act := Perimeter(Rectangle{10.0, 10.0})

	require.Equal(t, exp, act)
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, exp float64) {
		t.Helper()
		act := shape.Area()
		require.Equal(t, exp, act)
	}

	t.Run("rectangle", func(t *testing.T) {
		exp := 100.0
		rectangle := Rectangle{10.0, 10.0}

		checkArea(t, rectangle, exp)
	})
	t.Run("circle", func(t *testing.T) {
		exp := 314.1592653589793

		circle := Circle{10.0}

		checkArea(t, circle, exp)
	})
}

func Example_floatComparison() {
	fmt.Println(0.1+0.2 == 0.3)

	var a, b, c float64 = 0.1, 02, 0.3
	fmt.Println(a+b == c)
	// Output:
	// true
	// false
}
