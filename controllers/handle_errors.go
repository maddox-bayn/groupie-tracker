package controllers

import (
	"groupie-tracker/model"
	"log"
	"net/http"
)

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
