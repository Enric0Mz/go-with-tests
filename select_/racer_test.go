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
	slowServer := BuildDelayedServer(20 * time.Millisecond)
	fastServer := BuildDelayedServer(1 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	expect := fastServer.URL

	actual := Racer(slowServer.URL, fastServer.URL)

	require.Equal(t, expect, actual)

	slowServer.Close()
	fastServer.Close()
}
