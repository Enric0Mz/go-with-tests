package reflection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name   string
		Input  interface{}
		Expect []string
	}{
		{Name: "one string", Input: struct{ Name string }{"Lucia"}, Expect: []string{"Lucia"}},
		{Name: "two strings", Input: struct {
			Name string
			City string
		}{"Lucia", "Miami"}, Expect: []string{"Lucia", "Miami"}},
		{Name: "one string and one integer", Input: struct {
			Name string
			Age  int
		}{"Jason", 28}, Expect: []string{"Jason"}},
		{Name: "Nested struct", Input: Person{
			"Jason", Profile{
				City: "Vice City",
				Age:  28,
			},
		}, Expect: []string{"Jason", "Vice City"}},
		{Name: "pointers to things", Input: &Person{
			"Trevor", Profile{
				City: "Los Santos",
				Age:  47,
			},
		}, Expect: []string{"Trevor", "Los Santos"}},
		{Name: "slices", Input: []Profile{
			{33, "Sofia"},
			{66, "Athenas"},
		}, Expect: []string{"Sofia", "Athenas"}},
		{Name: "arrays", Input: [2]Profile{
			{1, "down"},
			{2, "up"},
		}, Expect: []string{"down", "up"}},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {
			var actual []string
			walk(test.Input, func(input string) {
				actual = append(actual, input)
			})
			require.Equal(t, test.Expect, actual)
		})
	}
	t.Run("with maps", func(t *testing.T) {
		mapCase := map[string]string{
			"Jason": "DuVal",
			"Lucia": "Caminos",
		}
		expect := []string{"DuVal", "Caminos"}
		var actual []string
		walk(mapCase, func(input string) {
			actual = append(actual, input)
		})
		require.Len(t, expect, len(actual))
		require.True(t, assertContains(t, actual, expect[0]))
		require.True(t, assertContains(t, actual, expect[1]))

	})
}

func assertContains(t testing.TB, arr []string, target string) bool {
	t.Helper()
	contains := false
	for _, curr := range arr {
		if curr == target {
			contains = true
		}
	}
	return contains
}
