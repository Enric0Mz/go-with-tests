package select_

import (
	"fmt"
	"net/http"
	"time"
)

const timeOut10Seconds = time.Second * 10

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, timeOut10Seconds)
}

func ConfigurableRacer(a, b string, timeOutDuration time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeOutDuration):
		return "", fmt.Errorf("Time out error")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			resp.Body.Close()
		}
		close(ch)
	}()
	return ch
}
