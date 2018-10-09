package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!")
}
func main() {
	hh := http.HandlerFunc(helloHandler)
	http.Handle("/", hh)
	http.ListenAndServe(":12345", nil)
}