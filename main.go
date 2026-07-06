package main

import (
	"fmt"
	cl "groupie-tracker/controllers"
	"html/template"
	"log"
	"net/http"
)

const url = "https://groupietrackers.herokuapp.com/api"
const port = ":8080"

// func FetchApi() error {
// 	Client := http.Client{Timeout: 10 * time.Second}
// 	http.Get(url)
// }

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}
	cl.RendersTemplates(w, http.StatusOK, "index.html", nil)
}
func RenderError(w http.ResponseWriter, code int, message string) {
	cl.RendersTemplates(w, code, "error.html", nil)
}
func errorHandler(w http.ResponseWriter, r *http.Request) {
	RenderError(w, http.StatusNotImplemented, "Error")
}
func main() {
	cl.Tmpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("GET /", HandleIndex)
	http.HandleFunc("GET /error", errorHandler)
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, nil))
}
