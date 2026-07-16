package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
// global template to parse all template
var Tmpl *template.Template

// function to preload all template at before start of the server
// to the global Template varaible and log if error occur
func ParseTemplates() {
	var err error
	Tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("error from parseGlob %s", err)
		log.Fatal(err)
	}
}

// RenderTemplate render a preloaded template with it data using Lookup method to get
// specific template to be renderd and return any internal server error
func RendersTemplates(w http.ResponseWriter, statuscode int, tmpl string, data any) error {
	template := Tmpl.Lookup(tmpl)
	if template == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return fmt.Errorf("error %s Not found", tmpl)
	}
	w.WriteHeader(statuscode)
	return template.Execute(w, data)
}
