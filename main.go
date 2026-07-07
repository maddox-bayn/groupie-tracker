package main

import (
	"fmt"
	"groupie-tracker/control_utils"
	cl "groupie-tracker/controllers"
	"groupie-tracker/data"
	"html/template"
	"log"
	"net/http"
)

const port = ":8080"

func init() {
	var err error
	data.CombinedData, err = control_utils.FtchAllData()
	if err != nil {
		log.Println("Failed to fetch data from api")
	}
	fmt.Println("successful data fetch")
}

func main() {
	if len(data.CombinedData.Date) == 0 {
		fmt.Println("FtchAllData failed to fetch data at init call.... retrying call again")
		var err error
		data.CombinedData, err = control_utils.FtchAllData(); if err != nil {
			log.Fatalf("Error fetching data:%v",err)
		}
	}
	cl.Tmpl = template.Must(template.ParseGlob("templates/*.html"))
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(port, nil))
}
