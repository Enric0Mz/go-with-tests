package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDonwStart = 3
const finalLine = "Go!"

func Countdown(writter io.Writer) {

	for i := countDonwStart; i > 0; i-- {
		fmt.Fprintln(writter, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(writter, finalLine)

}

func main() {
	Countdown(os.Stdout)
}
