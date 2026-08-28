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

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.shape.Area(), tt.want)
		})
	}

}

func Example_floatComparison() {
	fmt.Println(0.1+0.2 == 0.3)

	var a, b, c float64 = 0.1, 02, 0.3
	fmt.Println(a+b == c)
	// Output:
	// true
	// false
}
