package main

import (
	"fmt"
	"groupie-tracker/control_utils"
	cl "groupie-tracker/controllers"
	"groupie-tracker/data"
	"log"
	"net/http"
	"strings"
)

const port = ":8080"

func init() {
	var err error
	data.CombinedData, err = control_utils.FtchAllData()
	if err != nil {
		log.Println("Failed to fetch data from api")
	}
	if err == nil {
		fmt.Println("successful data fetch")
	}

}

func main() {
	if len(data.CombinedData.Dates) == 0 {
		fmt.Println("FtchAllData failed to fetch data at init call.... retrying call again")
		var err error
		data.CombinedData, err = control_utils.FtchAllData()
		if err != nil {
			log.Fatalf("Error fetching data:%v", err)
		}
	}
	http.HandleFunc("/static/", HandleStatic)
	cl.ParseTemplates()
	http.HandleFunc("/", cl.HandleMain)
	http.HandleFunc("/artist", cl.HandleArtist)
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, nil))
}


func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/static") {
		cl.RenderError(w, http.StatusUnauthorized, "the page does not exist.")
	}
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
}
