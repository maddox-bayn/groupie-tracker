package controllers

import (
	"groupie-tracker/model"
	"log"
	"net/http"
)

// RenderError function render  error page by calling RenderTemplate and filling it
// with the required data for the error page if any error occurs call http error and return
func RenderError(w http.ResponseWriter, statuscode int, message string) {
	err := RendersTemplates(w, statuscode, "error.html", model.Errors{
		Status:  statuscode,
		Message: message,
		Error:   http.StatusText(statuscode),
	})
	if err != nil {
		http.Error(w, "Internal Server Error", statuscode)
		log.Println("error from templaterendering")
		return
	}
}
