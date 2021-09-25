package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Writehellooooooo")
}

func (h hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I gotta get some screws")
}

func main() {
	var d hotdog
	var c hotcat
	http.Handle("/dog", d)
	http.Handle("/cat", c)
	http.ListenAndServe(":8080", nil)
}
