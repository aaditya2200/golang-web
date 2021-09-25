package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Something")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
