package slicesarrays

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	t.Run("collection of non fixed numbers", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 5, 6, 6, 7, 7}

		want := 46

		actual := Sum(slice)
		require.Equal(t, want, actual)
	})
}

func TestSumSlices(t *testing.T) {
	t.Run("collection of more than one slice", func(t *testing.T) {
		sliceOfSlices := [][]int{{1, 2}, {3, 4}}

		want := []int{3, 7}

		actual := SumSlices(sliceOfSlices)

		if !slices.Equal(actual, want) {
			t.Errorf("got %v want %v", actual, want)
		}
	})
}

func TestSumAllTaills(t *testing.T) {
	t.Run("run with some valid slices", func(t *testing.T) {

		sliceOfSlices := [][]int{{1, 2}, {3, 4}}
		want := []int{2, 4}

		actual := SumAllTaills(sliceOfSlices)

		require.Equal(t, want, actual)
	})
	t.Run("run with tottaly empty slice", func(t *testing.T) {
		emptySlice := [][]int{}

		want := []int{}
		actual := SumAllTaills(emptySlice)

		require.Equal(t, want, actual)
	})
	t.Run("run with some empty slices", func(t *testing.T) {
		someEmptySlices := [][]int{{}, {1, 2, 3}}

		want := []int{0, 5}
		actual := SumAllTaills(someEmptySlices)

		require.Equal(t, want, actual)
	})

}
