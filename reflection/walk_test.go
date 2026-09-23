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
}
