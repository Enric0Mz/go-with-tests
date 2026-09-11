package select_

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()

	resp, err := http.Get(a)

	if err == nil {
		resp.Body.Close()
	}
	timeA := time.Since(startA)
	fmt.Printf("Server A took %s miliseconds\n", timeA)

	startB := time.Now()

	resp, err = http.Get(b)

	if err != nil {
		resp.Body.Close()
	}

	timeB := time.Since(startB)
	fmt.Printf("Server B took %s miliseconds", timeB)

	if timeA > timeB {
		return b
	}
	return a

}
