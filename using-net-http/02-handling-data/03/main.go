package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("abc" , "bijkle")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<h1>Any code here</h1>`)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
