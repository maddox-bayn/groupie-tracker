package main

import (
	"fmt"
	"groupie-tracker/config"
	cu "groupie-tracker/control_utils"
	cl "groupie-tracker/controllers"
	"groupie-tracker/data"
	"log"
	"net/http"
)
// init function to load or fetch data from all api endpoint 
// before the the entry point of the program 
func init() {
	var err error
	data.CombinedData, err = cu.FtchAllData()
	if err != nil {
		log.Println("Failed to fetch data from api")
	}
	if err == nil {
		fmt.Println("successful data fetch")
	}

}

// main entery point of the  program and caller  
func main() {
	// if data was not loaded at init call, try fetch it again
	if len(data.CombinedData.Dates) == 0 {
		fmt.Println("FtchAllData failed to fetch data at init call.... retrying call again")
		var err error
		data.CombinedData, err = cu.FtchAllData()
		if err != nil {
			log.Fatalf("Error fetching data:%v", err)
		}
	}
	// function preload and store each artist data in a map[int]model.Artist 
	// for easy lookup of artist before system call
	cu.BuildArtistIndex()
	cu.BuildSearchIndex()

	// handle request for static pages using constum handler
	http.HandleFunc("/static/", cl.HandleStatic)

	// parse all template before start of the server
	cl.ParseTemplates()
	http.HandleFunc("/", cl.HandleMain)
	http.HandleFunc("/artist", cl.HandleArtist)
	// starting server
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(config.Port, nil))
}
