package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "Hello doggo")
	case "/cat":
		io.WriteString(w, "Hello kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
