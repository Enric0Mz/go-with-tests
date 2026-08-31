package maps

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDictionarySearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a test dict",
	}
	t.Run("Known word", func(t *testing.T) {
		expect := "this is a test dict"
		actual, err := dictionary.Search("test")
		require.NoError(t, err)
		require.Equal(t, expect, actual)

	})

	t.Run("Unknown word", func(t *testing.T) {
		expect := "word 'word2' not found in dictionary"
		_, err := dictionary.Search("word2")

		if !errors.Is(err, ErrNotFound) {
			panic("Error is not ErrNotFound")
		}

		require.Error(t, err)
		require.Equal(t, expect, err.Error())

	})
}

func TestDictionaryAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add("test", "this is a test word")

		require.NoError(t, err)
		expect := "this is a test word"
		assertDefinition(t, dictionary, "test", expect)

	})
	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test word"}

		err := dictionary.Add("test", "trying to update existing word")
		expect := "word already exists, aborting insertion"

		if !errors.Is(err, ErrConflict) {
			panic("error should be ErrConflict")
		}

		require.Error(t, err)
		require.Equal(t, expect, err.Error())
	})

}

func assertDefinition(t testing.TB, dict Dictionary, w string, expect string) {
	t.Helper()
	actual, err := dict.Search(w)

	require.NoError(t, err)
	require.Equal(t, expect, actual)
}
