package select_

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	timeA := measureResponseTime(a)
	timeB := measureResponseTime(b)

	if timeA > timeB {
		return b
	}
	return a

}

func measureResponseTime(url string) time.Duration {
	start := time.Now()

	resp, err := http.Get(url)

	if err == nil {
		resp.Body.Close()
	}
	timeA := time.Since(start)
	fmt.Printf("Server A took %s miliseconds\n", timeA)
	return timeA
}
