package controllers

import (
	"html/template"
	"log"
	"net/http"
)

var Tmpl *template.Template

func ParseTemplates() {
	var err error
	Tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("error from parseGlob %s", err)
	}
}

func RendersTemplates(w http.ResponseWriter, statuscode int, tmpl string, data any) {
	template := Tmpl.Lookup(tmpl)
	if template == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err := template.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
