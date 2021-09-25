package main

import (
	"io"
	"net/http"
)

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Helloooooooo")
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pottah")
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil)
}
