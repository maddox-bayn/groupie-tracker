package controllers

import (
	"fmt"
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
		log.Fatal(err)
	}
}

func RendersTemplates(w http.ResponseWriter, statuscode int, tmpl string, data any) error {	
	w.WriteHeader(statuscode)
	template := Tmpl.Lookup(tmpl)
	if template == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return fmt.Errorf("Erroo %s Not found", tmpl)
	}
	return template.Execute(w, data)
}
