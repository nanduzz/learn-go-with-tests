package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// it can only run in package main
func main() {
	Greet(os.Stdout, "Fernando")
	http.ListenAndServe(":5050", http.HandlerFunc(MyGreeterHandler))
}
