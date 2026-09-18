package select_

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func BuildDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("should return fastServer when less milisesconds than slowServer are passed", func(t *testing.T) {
		slowServer := BuildDelayedServer(20 * time.Millisecond)
		fastServer := BuildDelayedServer(1 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		expect := fastServer.URL

		actual, _ := Racer(slowServer.URL, fastServer.URL)

		require.Equal(t, expect, actual)
	})
	t.Run("should get timeout error when server's do not run in the specified time", func(t *testing.T) {
		timeOutOne := BuildDelayedServer(2 * time.Millisecond)
		timeOutTwo := BuildDelayedServer(2 * time.Millisecond)

		defer timeOutOne.Close()
		defer timeOutTwo.Close()

		_, err := ConfigurableRacer(timeOutOne.URL, timeOutTwo.URL, 1*time.Millisecond)

		require.Error(t, err)
	})

}
