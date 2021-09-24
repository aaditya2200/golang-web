package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Hello there")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
	fmt.Println(tpl.Templates())
}
