package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "something doggo")
}

func (h hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello catto")
}

func main() {
	var d hotdog
	var c hotcat
	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}
