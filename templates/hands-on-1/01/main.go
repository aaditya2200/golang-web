package main

import "text/template"

var tpl *template.Template

type hotels struct {
	Name string
	Address string
	City string
	zip int
	region string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

