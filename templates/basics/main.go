package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	name := "Aaditya Rangarajan"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	nf, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
}
