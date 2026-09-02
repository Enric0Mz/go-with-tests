package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writter io.Writer, n string) {
	fmt.Fprintf(writter, "Hello, %s", n)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
