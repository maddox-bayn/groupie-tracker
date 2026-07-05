package controllers

import (
	"html/template"
	"log"
)


var Tmpl *template.Template

func ParseTemplates()  {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("error from parseGlob %s", err)
	}
}