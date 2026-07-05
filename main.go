package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	cl "groupie-tracker/controllers"
)

const url = "https://groupietrackers.herokuapp.com/api"
const port = ":8080"

func FetchApi() error {
	Client := http.Client{Timeout: 10 * time.Second}
	http.Get(url)
}
func main() {
	cl.Tmpl = template.Must(template.ParseGlob("templates/*.html"))
	fs := http.FileServer(http.Dir("/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, nil))
}
