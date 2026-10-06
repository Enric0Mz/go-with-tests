package counter_

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCounter(t *testing.T) {
	t.Run("should increment 3 times when calling inc 3 times", func(t *testing.T) {
		cnt := Counter{}
		cnt.Inc()
		cnt.Inc()
		cnt.Inc()
		expect := 3
		actual := cnt.Size

		assert.Equal(t, expect, actual)
	})

	t.Run("should run safely when running concurrently", func(t *testing.T) {
		expectedCount := 1000
		cnt := Counter{}

		var wg sync.WaitGroup

		wg.Add(expectedCount)

		for range expectedCount {
			go func() {
				cnt.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		require.Equal(t, expectedCount, cnt.Size)
	})

	t.Run("should correctly increment with atomic Size", func(t *testing.T) {
		cnt := &CounterAtomic{}

		var wg sync.WaitGroup

		expectedCount := 100

		for range expectedCount {
			go func() {
				cnt.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		require.Equal(t, expectedCount, cnt.Value())
	})
}

func BenchmarkCounter(t *testing.B) {
	t.Run("Sync calls to Increment", func(b *testing.B) {
		counter := Counter{}
		for range 1000 {
			counter.Inc()
		}
	})

	t.Run("Async calls to increment with goroutines", func(b *testing.B) {
		counter := Counter{}
		expectedCount := 1000
		var wg sync.WaitGroup

		wg.Add(expectedCount)

		for range expectedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
