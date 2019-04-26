package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	FirstName string
	Age int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person {
		"David",
		37,
	}
	
	//tpl.ExecuteTemplate(w, "index.gohtml", nil)
	tpl.ExecuteTemplate(w, "index.gohtml", p1)
	
}